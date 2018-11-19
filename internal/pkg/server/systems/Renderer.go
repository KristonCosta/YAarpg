package systems

import (
	"fmt"

	"github.com/Notserc/go-pixel/internal/pkg/ecs"
	c "github.com/Notserc/go-pixel/internal/pkg/server/components"
)

type Renderer struct {
	World *ecs.World
	Types []ecs.ComponentType
}

func registerRenderer(world *ecs.World) {
	var renderer ecs.System
	renderer = &Renderer{
		World: world,
		Types: []ecs.ComponentType{c.PositionType, c.RenderableType},
	}
	world.AddSystem(&renderer)
}

func (self *Renderer) Update(dt float64) {
	for _, entity := range self.World.Entities {
		if self.World.HasComponents(entity, self.RequiredTypes()) {
			position := c.GetPosition(entity, self.World)
			renderable := c.GetRenderable(entity, self.World)
			fmt.Printf("Renderer: Entity %v has required components %+v %+v\n", entity, position, renderable)
		}
	}
}

func (self *Renderer) RequiredTypes() (types *[]ecs.ComponentType) {
	return &self.Types
}
