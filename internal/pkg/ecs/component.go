package ecs

type Component interface {
	Type() *ComponentType
}

type ComponentType = string

type ComponentManager struct {
	entities   []Entity
	components map[Entity]Component
}

func newManager() *ComponentManager {
	var manager ComponentManager
	components := make(map[Entity]Component)
	manager = ComponentManager{
		entities:   make([]Entity, 0),
		components: components}
	return &manager
}

func (manager *ComponentManager) hasComponent(entity Entity) bool {
	_, ok := manager.components[entity]
	return ok
}

func (manager *ComponentManager) getComponent(entity Entity) *Component {
	component, _ := manager.components[entity]
	return &component
}

func (manager *ComponentManager) addEntity(entity Entity, component Component) {
	manager.entities = append(manager.entities, entity)
	manager.components[entity] = component
}

func (manager *ComponentManager) deleteEntity(entity Entity) {
	delete(manager.components, entity)

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
