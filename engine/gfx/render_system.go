package gfx

import (
	"github.com/hajimehoshi/ebiten/v2"
	"rp-go-v2-physics-integrated/engine/ecs"
)

// RenderSystem handles drawing all sprites in the world using
// Transform and Camera data. Supports sprite scaling + camera zoom.
type RenderSystem struct {
	Sprites    ecs.ComponentMap[Sprite]
	Transforms ecs.ComponentMap[Transform]
	Cameras    ecs.ComponentMap[Camera]
}

// Draw renders all entities that have a Sprite and Transform component.
// Applies Sprite.Scale, Transform.Scale, and Camera.Zoom together.
func (rs *RenderSystem) Draw(screen *ebiten.Image) {
	// Determine active camera zoom (default = 1.0)
	camZoom := 1.0
	if cam, ok := rs.Cameras.First(); ok && cam.Zoom != 0 {
		camZoom = cam.Zoom
	}

	// Iterate through all entities with both Sprite + Transform
	for e := range rs.Sprites.Entities {
		sprite, okS := rs.Sprites.Get(e)
		transform, okT := rs.Transforms.Get(e)
		if !okS || !okT || sprite.Img == nil {
			continue
		}

		opts := &ebiten.DrawImageOptions{}
		bounds := sprite.Img.Bounds()
		w, h := float64(bounds.Dx()), float64(bounds.Dy())

		// Center origin before scaling
		opts.GeoM.Translate(-w/2, -h/2)

		// Combine scales
		totalScale := sprite.Scale * transform.Scale * camZoom
		if totalScale == 0 {
			totalScale = 1.0 // safety fallback
		}
		opts.GeoM.Scale(totalScale, totalScale)

		// Translate to world position
		opts.GeoM.Translate(transform.X, transform.Y)

		screen.DrawImage(sprite.Img, opts)
	}
}

// Update exists to satisfy the ECS System interface but isn’t used here.
func (rs *RenderSystem) Update(dt float64) {}

