package yaarpg

import (
	"fmt"
	"image"
	_ "image/png"
	"math/rand"
	"os"
	"time"

	"github.com/Notserc/go-pixel/internal/pkg/ecs"
	"github.com/Notserc/go-pixel/internal/pkg/server/components"

	"github.com/Notserc/go-pixel/internal/pkg/server"

	"github.com/faiface/pixel"

	"github.com/faiface/pixel/pixelgl"
	"github.com/pkg/profile"
	"golang.org/x/image/colornames"
)

func loadPicture(path string) (pixel.Picture, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}
	return pixel.PictureDataFromImage(img), nil
}

func run() {
	server := server.Init()
	cfg := pixelgl.WindowConfig{
		Title:  "Pixel Rocks!",
		Bounds: pixel.R(0, 0, 1024, 768),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}
	fps := time.Tick(time.Second / 120)
	renderTypes := []ecs.ComponentType{components.PositionType, components.RenderableType}
	win.Clear(colornames.Skyblue)

	spritesheet, err := loadPicture("assets/trees.png")
	if err != nil {
		panic(err)
	}
	var treesFrames []pixel.Rect
	for x := spritesheet.Bounds().Min.X; x < spritesheet.Bounds().Max.X; x += 32 {
		for y := spritesheet.Bounds().Min.Y; y < spritesheet.Bounds().Max.Y; y += 32 {
			treesFrames = append(treesFrames, pixel.R(x, y, x+32, y+32))
		}
	}
	var (
		frames = 0
		second = time.Tick(time.Second)
	)

	batch := pixel.NewBatch(&pixel.TrianglesData{}, spritesheet)
	_ = pixel.NewSprite(spritesheet, treesFrames[rand.Intn(len(treesFrames))])
	for !win.Closed() {

		server.Run()
		for _, entity := range server.World.Entities {
			if server.World.HasComponents(entity, &renderTypes) {
				// position := components.GetPosition(entity, server.World)
				// tree.Draw(batch, pixel.IM.Moved(pixel.V(position.X*512+512, position.Y*384+384)))

			}
		}
		frames++
		win.Clear(colornames.Aliceblue)
		batch.Draw(win)
		win.Update()
		batch.Clear()
		select {
		case <-second:
			win.SetTitle(fmt.Sprintf("%s | FPS: %d", cfg.Title, frames))
			frames = 0
		default:
		}
		<-fps
	}
}

func Run() {
	defer profile.Start().Stop()
	pixelgl.Run(run)
}
