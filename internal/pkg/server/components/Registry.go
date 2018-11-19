package components

import (
	"github.com/Notserc/go-pixel/internal/pkg/ecs"
)

func Register(world *ecs.World) {
	world.RegisterComponent(PositionType)
	world.RegisterComponent(RenderableType)
	world.RegisterComponent(AreaType)
	world.RegisterComponent(CollidableType)
	world.RegisterComponent(SpeedType)
}
