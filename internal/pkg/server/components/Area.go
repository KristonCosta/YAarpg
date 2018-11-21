package components

import (
	"github.com/Notserc/go-pixel/internal/pkg/ecs"
)

type Area struct {
	Left   float64
	Right  float64
	Top    float64
	Bottom float64
}

func (*Area) Type() *ecs.ComponentType {
	return &AreaType
}

func GetArea(entity ecs.Entity, world *ecs.World) *Area {
	return (*world.GetComponent(entity, AreaType)).(*Area)
}
