package gfx

import (
    "github.com/hajimehoshi/ebiten/v2"
    "rp-go-v2-physics-integrated/engine/ecs"
    "time"
)

type RenderSystem struct {
    Transforms *ecs.ComponentStore[Transform]
    Sprites    *ecs.ComponentStore[Sprite]
    Cameras    *ecs.ComponentStore[Camera]
}

func (r *RenderSystem) Init(w *ecs.World) {
    r.Transforms = ecs.NewComponentStore[Transform]()
    r.Sprites = ecs.NewComponentStore[Sprite]()
    r.Cameras = ecs.NewComponentStore[Camera]()
}

func (r *RenderSystem) Update(dt time.Duration) {}

func (r *RenderSystem) Draw(screen *ebiten.Image) {
    var cam Camera
    for _, c := range r.Cameras.All() {
        cam = c
        break
    }

    for id, sprite := range r.Sprites.All() {
        tr, ok := r.Transforms.Get(id)
        if !ok || sprite.Img == nil {
            continue
        }
        op := &ebiten.DrawImageOptions{}
        op.GeoM.Translate(-cam.X, -cam.Y)
        if cam.Zoom != 0 {
            op.GeoM.Scale(cam.Zoom, cam.Zoom)
        }
        op.GeoM.Translate(tr.X, tr.Y)
        op.GeoM.Scale(tr.Scale, tr.Scale)
        op.GeoM.Rotate(tr.Rotation)
        screen.DrawImage(sprite.Img, op)
    }
}

func (r *RenderSystem) Shutdown() {}

