package components

import (
	"github.com/Notserc/go-pixel/internal/pkg/ecs"
)

var AreaType ecs.ComponentType = 0
var CollidableType ecs.ComponentType = 1
var PositionType ecs.ComponentType = 2
var RenderableType ecs.ComponentType = 3
var SpeedType ecs.ComponentType = 4

func Register(world *ecs.World) {
	world.RegisterComponent(PositionType)
	world.RegisterComponent(RenderableType)
	world.RegisterComponent(AreaType)
	world.RegisterComponent(CollidableType)
	world.RegisterComponent(SpeedType)
}
