package event

import "sync"

type Event struct {
    Type    string
    Payload any
}

type Handler func(Event)

type EventBus struct {
    subs map[string][]Handler
    mu   sync.RWMutex
}

func NewEventBus() *EventBus {
    return &EventBus{subs: make(map[string][]Handler)}
}

func (b *EventBus) Subscribe(eventType string, handler Handler) {
    b.mu.Lock()
    defer b.mu.Unlock()
    b.subs[eventType] = append(b.subs[eventType], handler)
}

func (b *EventBus) Publish(ev Event) {
    b.mu.RLock()
    defer b.mu.RUnlock()
    if handlers, ok := b.subs[ev.Type]; ok {
        for _, h := range handlers {
            h(ev)
        }
    }
}

