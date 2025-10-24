package ecs

import (
    "time"
    "github.com/hajimehoshi/ebiten/v2"
)

type World struct {
    entities  *EntityManager
    systems   []System
    renderers []RenderSystem
}

func NewWorld() *World {
    return &World{entities: NewEntityManager()}
}

func (w *World) NewEntity() EntityID {
    return w.entities.NewEntity()
}

func (w *World) AddSystem(s System) {
    w.systems = append(w.systems, s)
    s.Init(w)
    if rs, ok := s.(RenderSystem); ok {
        w.renderers = append(w.renderers, rs)
    }
}

func (w *World) Update(dt time.Duration) {
    for _, s := range w.systems {
        s.Update(dt)
    }
}

func (w *World) Draw(screen *ebiten.Image) {
    for _, r := range w.renderers {
        r.Draw(screen)
    }
}

func (w *World) Shutdown() {
    for _, s := range w.systems {
        s.Shutdown()
    }
}
