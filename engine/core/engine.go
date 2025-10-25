package core

import (
	"github.com/hajimehoshi/ebiten/v2"
	"rp-go-v2-physics-integrated/engine/ecs"
    	"rp-go-v2-physics-integrated/engine/gfx"
    	"rp-go-v2-physics-integrated/engine/physics"
)

// Engine is the central orchestrator for all systems and the ECS world.
type Engine struct {
	world         *ecs.World
	renderSystem  *gfx.RenderSystem
	physicsSystem *physics.MovementSystem
}

// NewEngine initializes the ECS world and registers core systems.
func NewEngine() *Engine {
	world := ecs.NewWorld()

	engine := &Engine{
		world: world,
	}

	// Initialize and register systems
	engine.physicsSystem = physics.NewMovementSystem(world)
	engine.renderSystem = gfx.NewRenderSystem(world)

	world.RegisterSystem(engine.physicsSystem)
	world.RegisterSystem(engine.renderSystem)

	return engine
}

// Update advances the game state — called once per frame by Ebiten.
func (e *Engine) Update() error {
	// Update all ECS systems (physics, logic, etc.)
	e.world.Update(1.0 / 60.0) // fixed timestep for simplicity
	return nil
}

// Draw renders the current world state to the Ebiten screen.
func (e *Engine) Draw(screen *ebiten.Image) {
	e.renderSystem.Draw(screen)
}

// Layout defines the logical game resolution.
func (e *Engine) Layout(outsideWidth, outsideHeight int) (int, int) {
	// Adjust these to your target virtual resolution
	return 800, 600
}

