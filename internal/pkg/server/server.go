package server

import (
	"math/rand"
	"time"

	"github.com/Notserc/go-pixel/internal/pkg/ecs"
	"github.com/Notserc/go-pixel/internal/pkg/server/components"
	"github.com/Notserc/go-pixel/internal/pkg/server/systems"
)

type Server struct {
	World      *ecs.World
	lastUpdate time.Time
}

func createWorld() (world *ecs.World) {
	world = ecs.NewWorld()
	components.Register(world)
	systems.Register(world)
	return
}

func createBall(world *ecs.World) {
	ball := world.AddEntity()
	position := components.Position{
		X: 0.0,
		Y: 0.0,
	}
	speed := components.Speed{
		VX: 2*rand.Float64() - 1,
		VY: 2*rand.Float64() - 1,
	}
	collidable := components.Collidable{
		Radius: 0.05,
	}
	renderable := components.Renderable{
		Char: '1',
	}
	world.AddComponent(ball, &position)
	world.AddComponent(ball, &speed)
	world.AddComponent(ball, &collidable)
	world.AddComponent(ball, &renderable)
}

func Init() *Server {
	server := Server{
		World:      createWorld(),
		lastUpdate: time.Now()}
	createArea(server.World)
	for i := 1; i < 10000; i++ {
		createBall(server.World)
		createBall(server.World)
	}
	return &server
}

func createArea(world *ecs.World) {
	area := world.AddEntity()
	areaComp := components.Area{
		Left:   -1.0,
		Right:  1.0,
		Top:    1.0,
		Bottom: -1.0,
	}
	world.AddComponent(area, &areaComp)
}

func (server *Server) Run() {
	dt := time.Since(server.lastUpdate).Seconds()
	server.World.Update(dt)
	server.lastUpdate = time.Now()
}
