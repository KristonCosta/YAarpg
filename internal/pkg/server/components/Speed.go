package components

import (
	"github.com/Notserc/go-pixel/internal/pkg/ecs"
)

type Speed struct {
	VX float64
	VY float64
}

func (*Speed) Type() *ecs.ComponentType {
	return &SpeedType
}

func GetSpeed(entity ecs.Entity, world *ecs.World) *Speed {
	return (*world.GetComponent(entity, SpeedType)).(*Speed)
}
