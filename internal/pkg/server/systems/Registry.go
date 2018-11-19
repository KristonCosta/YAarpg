package systems

import (
	"github.com/Notserc/go-pixel/internal/pkg/ecs"
)

func Register(world *ecs.World) {
	registerSystemCollide(world)
	registerSystemDraw(world)
	registerSystemMove(world)
}
