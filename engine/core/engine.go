package core

import (
    "time"

    "github.com/hajimehoshi/ebiten/v2"
    "rp-go-v2-physics-integrated/engine/ecs"
    "rp-go-v2-physics-integrated/engine/event"
)

type Engine struct {
    bus        *event.EventBus
    systems    []System
    frameDelta time.Duration
    World      *ecs.World
}

type System interface {
    Init(*Engine)
    Update(dt time.Duration)
    Shutdown()
}

func New() *Engine {
    return &Engine{
        bus:   event.NewEventBus(),
        World: ecs.NewWorld(),
    }
}

func (e *Engine) AddSystem(sys System) {
    e.systems = append(e.systems, sys)
    sys.Init(e)
}

func (e *Engine) RunFrame(dt time.Duration) {
    e.frameDelta = dt
    for _, sys := range e.systems {
        sys.Update(dt)
    }
    if e.World != nil {
        e.World.Update(dt)
    }
}

func (e *Engine) Draw(screen *ebiten.Image) {
    if e.World != nil {
        e.World.Draw(screen)
    }
}

func (e *Engine) Shutdown() {
    for _, sys := range e.systems {
        sys.Shutdown()
    }
    if e.World != nil {
        e.World.Shutdown()
    }
}

func (e *Engine) Bus() *event.EventBus {
    return e.bus
}
