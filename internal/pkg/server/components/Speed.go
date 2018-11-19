package components

import (
	"github.com/Notserc/go-pixel/internal/pkg/ecs"
)

type Speed struct {
	VX float64
	VY float64
}

var SpeedType ecs.ComponentType = "speed"

func (*Speed) Type() *ecs.ComponentType {
	return &SpeedType
}

func GetSpeed(entity ecs.Entity, world *ecs.World) *Speed {
	return (*world.GetComponent(entity, SpeedType)).(*Speed)
}
