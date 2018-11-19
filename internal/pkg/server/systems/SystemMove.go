package systems

import (
	"github.com/Notserc/go-pixel/internal/pkg/ecs"
	c "github.com/Notserc/go-pixel/internal/pkg/server/components"
)

type SystemMove struct {
	World *ecs.World
	Types []ecs.ComponentType
}

func registerSystemMove(world *ecs.World) {
	var move ecs.System
	move = &SystemMove{
		World: world,
		Types: []ecs.ComponentType{c.PositionType, c.SpeedType},
	}
	world.AddSystem(&move)
}

func (self *SystemMove) Update(dt float64) {
	for _, entity := range self.World.Entities {
		if self.World.HasComponents(entity, self.RequiredTypes()) {
			position := c.GetPosition(entity, self.World)
			speed := c.GetSpeed(entity, self.World)
			position.X += (speed.VX) * dt
			position.Y += (speed.VY) * dt
		}
	}
}

func (self *SystemMove) RequiredTypes() (types *[]ecs.ComponentType) {
	return &self.Types
}
