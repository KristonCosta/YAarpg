package systems

import (
	"fmt"

	"github.com/Notserc/go-pixel/internal/pkg/ecs"
	c "github.com/Notserc/go-pixel/internal/pkg/server/components"
)

type Controller struct {
	World *ecs.World
	Types []ecs.ComponentType
}

func registerController(world *ecs.World) {
	var controller ecs.System
	controller = &Controller{
		World: world,
		Types: []ecs.ComponentType{c.PositionType},
	}
	world.AddSystem(&controller)
}

func (controller *Controller) Update(dt float64) {
	for _, entity := range controller.World.Entities {
		if controller.World.HasComponents(entity, controller.RequiredTypes()) {
			position := c.GetPosition(entity, controller.World)
			fmt.Printf("Controller: Entity %v has required components %+v\n", entity, position)
		}
	}
}

func (controller *Controller) RequiredTypes() (types *[]ecs.ComponentType) {
	return &controller.Types
}
