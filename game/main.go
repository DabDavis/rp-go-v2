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
	// Initialize core engine and ECS world
	eng := core.New()

	// Create and register systems
	renderSys := &gfx.RenderSystem{}
	moveSys := &physics.MovementSystem{}

	eng.World.AddSystem(renderSys)
	eng.World.AddSystem(moveSys)

	// === Entity: Player ship ===
	ship := eng.World.NewEntity()

	img, err := gfx.Assets.LoadImage("game/assets/ship.png")
	if err != nil {
		fmt.Println("Failed to load ship.png:", err)
		return
	}

	// Add sprite with scaling (0.2 = 20% of original texture)
	renderSys.Sprites.Add(ship, gfx.Sprite{
		Img:   img,
		Scale: 0.2,
	})

	// Transform defines world position and scale multiplier (for animations or effects)
	renderSys.Transforms.Add(ship, gfx.Transform{
		X:     320,
		Y:     240,
		Scale: 1.0, // can be animated dynamically
	})

	// Physics components
	moveSys.Transforms.Add(ship, gfx.Transform{X: 320, Y: 240, Scale: 1.0})
	moveSys.Velocities.Add(ship, physics.Velocity{Mode: physics.Momentum})

	// === Camera entity ===
	camEntity := eng.World.NewEntity()
	renderSys.Cameras.Add(camEntity, gfx.Camera{
		Zoom: 1.0, // default zoom factor (1.0 = 100%)
	})
	moveSys.Cameras.Add(camEntity, gfx.Camera{Zoom: 1.0})
	moveSys.Follows.Add(camEntity, physics.CameraFollow{Target: ship})

	// === Platform setup ===
	ebitenPlatform := platform.NewEbitenPlatform(eng, 800, 600)
	ebitenPlatform.Run(func(dt time.Duration) bool {
		eng.RunFrame(dt)
		return true
	})
}

