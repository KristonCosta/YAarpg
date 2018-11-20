package systems

import (
	"github.com/Notserc/go-pixel/internal/pkg/ecs"
	c "github.com/Notserc/go-pixel/internal/pkg/server/components"
)

type SystemDraw struct {
	World *ecs.World
	Types []ecs.ComponentType
}

func registerSystemDraw(world *ecs.World) {
	var draw ecs.System
	draw = &SystemDraw{
		World: world,
		Types: []ecs.ComponentType{c.PositionType},
	}
	world.AddSystem(&draw)
}

func (self *SystemDraw) Update(dt float64) {
	for _, entity := range self.World.Entities {
		if self.World.HasComponents(entity, self.RequiredTypes()) {
			//	position := c.GetPosition(entity, self.World)
			//	fmt.Printf("Entity %v (%v, %v)\n", entity, position.X, position.Y)
		}
	}
}

func (self *SystemDraw) RequiredTypes() (types *[]ecs.ComponentType) {
	return &self.Types
}
