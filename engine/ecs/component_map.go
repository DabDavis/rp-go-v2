package ecs

// ComponentMap is a lightweight generic container for ECS components.
type ComponentMap[T any] struct {
	Entities map[int]T
}

// Add registers or replaces a component for an entity ID.
func (m *ComponentMap[T]) Add(id int, comp T) {
	if m.Entities == nil {
		m.Entities = make(map[int]T)
	}
	m.Entities[id] = comp
}

// Get retrieves a component for an entity ID.
func (m *ComponentMap[T]) Get(id int) (T, bool) {
	val, ok := m.Entities[id]
	return val, ok
}

// First returns the first component (useful for cameras or singletons).
func (m *ComponentMap[T]) First() (T, bool) {
	for _, v := range m.Entities {
		return v, true
	}
	var zero T
	return zero, false
}

