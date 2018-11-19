package yaarpg

import (
	"math/rand"

	"github.com/Notserc/go-pixel/internal/pkg/ecs"
	"github.com/Notserc/go-pixel/internal/pkg/server/components"
	"github.com/Notserc/go-pixel/internal/pkg/server/systems"
)

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
		VX: 10.0 * rand.Float64(),
		VY: 10.0 * rand.Float64(),
	}
	collidable := components.Collidable{
		Radius: 0.05,
	}
	world.AddComponent(ball, &position)
	world.AddComponent(ball, &speed)
	world.AddComponent(ball, &collidable)
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

func Run() {
	world := createWorld()
	createArea(world)
	createBall(world)
	createBall(world)
	for i := 0; i < 20; i++ {
		world.Update(0.032)
	}
}
