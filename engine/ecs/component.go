package ecs

import "sync"

type ComponentStore[T any] struct {
    data map[EntityID]T
    mu   sync.RWMutex
}

func NewComponentStore[T any]() *ComponentStore[T] {
    return &ComponentStore[T]{data: make(map[EntityID]T)}
}

func (s *ComponentStore[T]) Add(e EntityID, c T) {
    s.mu.Lock()
    defer s.mu.Unlock()
    s.data[e] = c
}

func (s *ComponentStore[T]) Get(e EntityID) (T, bool) {
    s.mu.RLock()
    defer s.mu.RUnlock()
    c, ok := s.data[e]
    return c, ok
}

func (s *ComponentStore[T]) All() map[EntityID]T {
    s.mu.RLock()
    defer s.mu.RUnlock()
    out := make(map[EntityID]T, len(s.data))
    for k, v := range s.data {
        out[k] = v
    }
    return out
}
