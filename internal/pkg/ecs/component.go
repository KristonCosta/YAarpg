package ecs

type Component interface {
	Type() *ComponentType
}

type ComponentType = int

type ComponentManager struct {
	entities   []Entity
	components []Component
}

func newManager() *ComponentManager {
	var manager ComponentManager
	components := make([]Component, 500000)
	manager = ComponentManager{
		entities:   make([]Entity, 500000),
		components: components}
	return &manager
}

func (manager *ComponentManager) hasComponent(entity Entity) bool {
	return manager.components[entity] != nil
}

func (manager *ComponentManager) getComponent(entity Entity) *Component {
	component := manager.components[entity]
	return &component
}

func (manager *ComponentManager) addEntity(entity Entity, component Component) {
	manager.components[entity] = component
}

func (manager *ComponentManager) deleteEntity(entity Entity) {
	// delete(manager.components, entity)

	manager.components[entity] = nil
	println(entity)
	return
	var index int
	for i, v := range manager.entities {
		if v == entity {
			index = i
			break
		}
	}
	lastIndex := len(manager.entities) - 1
	if lastIndex < 0 {
		manager.entities = manager.entities[:0]
		return
	}
	manager.entities[index] = manager.entities[lastIndex]
	manager.entities = manager.entities[:lastIndex]
}
