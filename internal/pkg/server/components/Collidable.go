package components

import (
	"github.com/Notserc/go-pixel/internal/pkg/ecs"
)

type Collidable struct {
	Radius float64
}

var CollidableType ecs.ComponentType = "collidable"

func (*Collidable) Type() *ecs.ComponentType {
	return &CollidableType
}

func GetCollidable(entity ecs.Entity, world *ecs.World) *Collidable {
	return (*world.GetComponent(entity, CollidableType)).(*Collidable)
}
