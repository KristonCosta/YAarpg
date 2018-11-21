package components

import "github.com/Notserc/go-pixel/internal/pkg/ecs"

type Renderable struct {
	Char rune
}

func (self *Renderable) Type() *ecs.ComponentType {
	return &RenderableType
}

func GetRenderable(entity ecs.Entity, world *ecs.World) *Renderable {
	return (*world.GetComponent(entity, RenderableType)).(*Renderable)
}
