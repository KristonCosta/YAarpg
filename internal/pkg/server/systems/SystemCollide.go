package systems

import (
	"github.com/Notserc/go-pixel/internal/pkg/ecs"
	c "github.com/Notserc/go-pixel/internal/pkg/server/components"
	"github.com/SolarLune/resolv/resolv"
)

type SystemCollide struct {
	World *ecs.World
	Types []ecs.ComponentType
	space *resolv.Space
}

func registerSystemCollide(world *ecs.World) {
	var collide ecs.System
	collide = &SystemCollide{
		World: world,
		Types: []ecs.ComponentType{c.PositionType, c.CollidableType, c.SpeedType},
		space: resolv.NewSpace(),
	}
	world.AddSystem(&collide)
}

func (self *SystemCollide) Update(dt float64) {
	for _, entity := range self.World.Entities {
		if self.World.HasComponents(entity, self.RequiredTypes()) {
			position := c.GetPosition(entity, self.World)
			speed := c.GetSpeed(entity, self.World)
			// collidable := c.GetCollidable(entity, self.World)
			area := c.GetArea((*self.World.GetEntitiesWithComponent(c.AreaType))[0], self.World)
			if position.X >= area.Right {
				position.X -= (position.X - area.Right)
				speed.VX = -speed.VX
				//	fmt.Printf("Entity %v  Collision(right): (%v, %v)\n", entity, position.X, position.Y)
			} else if position.X <= area.Left {
				position.X += (area.Left - position.X)
				speed.VX = -speed.VX
				//	fmt.Printf("Entity %v  Collision(left): (%v, %v)\n", entity, position.X, position.Y)
			}
			if position.Y >= area.Top {
				position.Y -= (position.Y - area.Top)
				speed.VY = -speed.VY
				//	fmt.Printf("Entity %v  Collision(top): (%v, %v)\n", entity, position.X, position.Y)
			} else if position.Y <= area.Bottom {
				position.Y += (area.Bottom - position.Y)
				speed.VY = -speed.VY
				//	fmt.Printf("Entity %v  Collision(bottom): (%v, %v)\n", entity, position.X, position.Y)
			}

		}
	}
}

func (self *SystemCollide) RequiredTypes() (types *[]ecs.ComponentType) {
	return &self.Types
}
