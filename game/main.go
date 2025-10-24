package main

import (
    "fmt"
    "time"
    "rp-go-v2-physics-integrated/engine/core"
    "rp-go-v2-physics-integrated/engine/gfx"
    "rp-go-v2-physics-integrated/engine/platform"
    "rp-go-v2-physics-integrated/engine/physics"
)

func main() {
    eng := core.New()

    renderSys := &gfx.RenderSystem{}
    eng.World.AddSystem(renderSys)

    moveSys := &physics.MovementSystem{}
    eng.World.AddSystem(moveSys)

    ship := eng.World.NewEntity()
    img, err := gfx.Assets.LoadImage("game/assets/ship.png")
    if err != nil {
        fmt.Println("Failed to load ship.png:", err)
        return
    }
    renderSys.Sprites.Add(ship, gfx.Sprite{Img: img})
    renderSys.Transforms.Add(ship, gfx.Transform{X: 320, Y: 240, Scale: 1.0})

    moveSys.Transforms.Add(ship, gfx.Transform{X: 320, Y: 240, Scale: 1.0})
    moveSys.Velocities.Add(ship, physics.Velocity{Mode: physics.Momentum})

    camEntity := eng.World.NewEntity()
    renderSys.Cameras.Add(camEntity, gfx.Camera{Zoom: 1.0})
    moveSys.Cameras.Add(camEntity, gfx.Camera{Zoom: 1.0})
    moveSys.Follows.Add(camEntity, physics.CameraFollow{Target: ship})

    ebitenPlatform := platform.NewEbitenPlatform(eng, 800, 600)
    ebitenPlatform.Run(func(dt time.Duration) bool {
        eng.RunFrame(dt)
        return true
    })
}
