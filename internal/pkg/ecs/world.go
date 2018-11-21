package ecs

type emptyVal struct{}

type World struct {
	nextEntity Entity
	Entities   []Entity
	managers   []*ComponentManager
	systems    []*System
}

func NewWorld() *World {
	return &World{
		nextEntity: 0,
		Entities:   make([]Entity, 0, 10000),
		managers:   make([]*ComponentManager, 100),
		systems:    make([]*System, 0)}
}

func (world *World) RegisterComponent(componentType ComponentType) {
	world.managers[componentType] = newManager()
}

func (world *World) AddEntity() Entity {
	world.Entities = append(world.Entities, world.nextEntity)
	entity := world.nextEntity
	world.nextEntity++
	return entity
}

func (world *World) RemoveEntity(entity Entity) {
	for index, e := range world.Entities {
		if e == entity {
			world.Entities = append(world.Entities[:index], world.Entities[index+1:]...)
			break
		}
	}
	for _, manager := range world.managers {
		if manager != nil {
			manager.deleteEntity(entity)
		}
	}
}

func (world *World) GetEntitiesWithComponent(componentType ComponentType) *[]Entity {
	return &world.managers[componentType].entities
}

func (world *World) AddComponent(entity Entity, component Component) {
	world.managers[*component.Type()].addEntity(entity, component)
}

func (world *World) HasComponents(entity Entity, components *[]ComponentType) bool {
	for _, componentType := range *components {
		if !world.managers[componentType].hasComponent(entity) {
			return false
		}
	}
	return true
}

func (world *World) AddSystem(system *System) {
	world.systems = append(world.systems, system)
}

func (world *World) GetComponent(entity Entity, componentType ComponentType) (component *Component) {
	return world.managers[componentType].getComponent(entity)
}

func (world *World) Update(dt float64) {
	for _, system := range world.systems {
		(*system).Update(dt)
	}
}
