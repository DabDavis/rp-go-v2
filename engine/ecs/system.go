package ecs

import (
    "time"
    "github.com/hajimehoshi/ebiten/v2"
)

type System interface {
    Init(*World)
    Update(dt time.Duration)
    Shutdown()
}

type RenderSystem interface {
    System
    Draw(screen *ebiten.Image)
}
