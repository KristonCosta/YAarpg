package components

import (
	"github.com/Notserc/go-pixel/internal/pkg/ecs"
)

type Position struct {
	X float64
	Y float64
}

var PositionType ecs.ComponentType = "position"

func (*Position) Type() *ecs.ComponentType {
	return &PositionType
}

func GetPosition(entity ecs.Entity, world *ecs.World) *Position {
	return (*world.GetComponent(entity, PositionType)).(*Position)
}
