package systems

import (
	"fmt"

	"github.com/Notserc/go-pixel/internal/pkg/ecs"
	c "github.com/Notserc/go-pixel/internal/pkg/server/components"
)

type GoldSystem struct {
	World *ecs.World
	Types []ecs.ComponentType
}

func registerGoldSystem(world *ecs.World) {
	var goldSystem ecs.System
	goldSystem = &GoldSystem{
		World: world,
		Types: []ecs.ComponentType{c.GoldType},
	}
	world.AddSystem(&goldSystem)
}

func (self *GoldSystem) Update(dt float64) {
	for _, entity := range self.World.Entities {
		if self.World.HasComponents(entity, self.RequiredTypes()) {
			gold := c.GetGold(entity, self.World)
			gold.GoldAmount = gold.Generation + gold.GoldAmount
			fmt.Printf("Gold: Entity %v has required components %+v\n", entity, gold)
		}
	}
}

func (self *GoldSystem) RequiredTypes() (types *[]ecs.ComponentType) {
	return &self.Types
}
