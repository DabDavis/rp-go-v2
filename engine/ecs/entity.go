package ecs

import "sync"

type EntityID uint64

type EntityManager struct {
    nextID   EntityID
    recycled []EntityID
    mu       sync.Mutex
}

func NewEntityManager() *EntityManager { return &EntityManager{} }

func (m *EntityManager) NewEntity() EntityID {
    m.mu.Lock()
    defer m.mu.Unlock()
    if len(m.recycled) > 0 {
        id := m.recycled[len(m.recycled)-1]
        m.recycled = m.recycled[:len(m.recycled)-1]
        return id
    }
    m.nextID++
    return m.nextID
}

func (m *EntityManager) DestroyEntity(id EntityID) {
    m.mu.Lock()
    defer m.mu.Unlock()
    m.recycled = append(m.recycled, id)
}
