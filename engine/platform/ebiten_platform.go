package platform

import (
    "log"
    "time"

    "github.com/hajimehoshi/ebiten/v2"
    "rp-go-v2-physics-integrated/engine/core"
)

type EbitenPlatform struct {
    engine     *core.Engine
    lastUpdate time.Time
}

func NewEbitenPlatform(engine *core.Engine, w, h int) *EbitenPlatform {
    ebiten.SetWindowSize(w, h)
    ebiten.SetWindowTitle("rp-go v2 Physics Integration")
    ebiten.SetWindowResizable(true)
    return &EbitenPlatform{engine: engine, lastUpdate: time.Now()}
}

func (p *EbitenPlatform) Run(mainLoop func(dt time.Duration) bool) {
    if err := ebiten.RunGame(p); err != nil {
        log.Fatal(err)
    }
}

func (p *EbitenPlatform) Update() error {
    now := time.Now()
    dt := now.Sub(p.lastUpdate)
    p.lastUpdate = now
    p.engine.RunFrame(dt)
    return nil
}

func (p *EbitenPlatform) Draw(screen *ebiten.Image) {
    p.engine.Draw(screen)
}

func (p *EbitenPlatform) Layout(outW, outH int) (int, int) { return outW, outH }
