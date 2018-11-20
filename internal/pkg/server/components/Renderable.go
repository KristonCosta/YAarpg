package components

import "github.com/Notserc/go-pixel/internal/pkg/ecs"

type Renderable struct {
	Char rune
}

var RenderableType ecs.ComponentType = 3

func (self *Renderable) Type() *ecs.ComponentType {
	return &RenderableType
}

func GetRenderable(entity ecs.Entity, world *ecs.World) *Renderable {
	return (*world.GetComponent(entity, RenderableType)).(*Renderable)
}
