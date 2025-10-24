package physics

import (
    "math"
    "github.com/hajimehoshi/ebiten/v2"
    "rp-go-v2-physics-integrated/engine/ecs"
    "rp-go-v2-physics-integrated/engine/gfx"
    "time"
)

type MovementMode int

const (
    Arcade MovementMode = iota
    Momentum
)

type Velocity struct {
    VX, VY        float64
    RotationSpeed float64
    Mode          MovementMode
}

type CameraFollow struct{ Target ecs.EntityID }

type MovementSystem struct {
    Transforms *ecs.ComponentStore[gfx.Transform]
    Velocities *ecs.ComponentStore[Velocity]
    Cameras    *ecs.ComponentStore[gfx.Camera]
    Follows    *ecs.ComponentStore[CameraFollow]
}

const (
    Accel     = 150.0
    Friction  = 0.98
    MaxSpeed  = 600.0
    RotSpeed  = 2.5
    ZoomStep  = 0.1
    ZoomMin   = 0.5
    ZoomMax   = 3.0
)

func (m *MovementSystem) Init(w *ecs.World) {
    m.Transforms = ecs.NewComponentStore[gfx.Transform]()
    m.Velocities = ecs.NewComponentStore[Velocity]()
    m.Cameras = ecs.NewComponentStore[gfx.Camera]()
    m.Follows = ecs.NewComponentStore[CameraFollow]()
}

func (m *MovementSystem) Update(dt time.Duration) {
    seconds := dt.Seconds()

    // --- Handle movement ---
    for id, v := range m.Velocities.All() {
        tr, ok := m.Transforms.Get(id)
        if !ok {
            continue
        }

        switch v.Mode {
        case Arcade:
            if ebiten.IsKeyPressed(ebiten.KeyW) || ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
                tr.Y -= Accel * seconds
            }
            if ebiten.IsKeyPressed(ebiten.KeyS) || ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
                tr.Y += Accel * seconds
            }
            if ebiten.IsKeyPressed(ebiten.KeyA) || ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
                tr.X -= Accel * seconds
            }
            if ebiten.IsKeyPressed(ebiten.KeyD) || ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
                tr.X += Accel * seconds
            }

        case Momentum:
            thrust := 0.0
            if ebiten.IsKeyPressed(ebiten.KeyW) || ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
                thrust = Accel
            } else if ebiten.IsKeyPressed(ebiten.KeyS) || ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
                thrust = -Accel / 2
            }

            if ebiten.IsKeyPressed(ebiten.KeyA) || ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
                tr.Rotation -= RotSpeed * seconds
            }
            if ebiten.IsKeyPressed(ebiten.KeyD) || ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
                tr.Rotation += RotSpeed * seconds
            }

            // Apply acceleration based on rotation
            v.VX += math.Sin(tr.Rotation) * thrust * seconds
            v.VY -= math.Cos(tr.Rotation) * thrust * seconds

            // Apply friction
            v.VX *= Friction
            v.VY *= Friction

            // Clamp speed
            spd := math.Sqrt(v.VX*v.VX + v.VY*v.VY)
            if spd > MaxSpeed {
                scale := MaxSpeed / spd
                v.VX *= scale
                v.VY *= scale
            }

            tr.X += v.VX * seconds
            tr.Y += v.VY * seconds
        }

        m.Transforms.Add(id, tr)
        m.Velocities.Add(id, v)
    }

    // --- Handle camera follow ---
    for _, follow := range m.Follows.All() {
        camEntity := follow.Target
        for _, cam := range m.Cameras.All() {
            target, ok := m.Transforms.Get(camEntity)
            if !ok {
                continue
            }
            cam.X += (target.X - cam.X) * 0.1
            cam.Y += (target.Y - cam.Y) * 0.1

            // Zoom control
            if ebiten.IsKeyPressed(ebiten.KeyQ) {
                cam.Zoom -= ZoomStep * seconds * 10
            }
            if ebiten.IsKeyPressed(ebiten.KeyE) {
                cam.Zoom += ZoomStep * seconds * 10
            }
            if cam.Zoom < ZoomMin {
                cam.Zoom = ZoomMin
            }
            if cam.Zoom > ZoomMax {
                cam.Zoom = ZoomMax
            }

            m.Cameras.Add(camEntity, cam)
        }
    }
}

func (m *MovementSystem) Shutdown() {}

