package gfx

import (
	"github.com/hajimehoshi/ebiten/v2"
)

// RenderSystem draws all visible sprites, taking into account
// per-sprite scale, per-entity transform, and active camera zoom.
type RenderSystem struct {
	Sprites    ComponentMap[Sprite]
	Transforms ComponentMap[Transform]
	Cameras    ComponentMap[Camera]
}

// Sprite is a renderable image with its own local scale.
type Sprite struct {
	Img   *ebiten.Image
	Scale float64 // 1.0 = full size, 0.2 = 20%, etc.
}

// Transform defines position and optional scale multiplier.
type Transform struct {
	X, Y  float64
	Scale float64 // 1.0 = normal; can animate separately from sprite scale
}

// Camera controls global zoom and could later support offsets.
type Camera struct {
	Zoom float64 // 1.0 = 100% zoom, 0.5 = zoom out, 2.0 = zoom in
}

// Draw renders all sprites to the screen, applying scale and camera zoom.
func (rs *RenderSystem) Draw(screen *ebiten.Image) {
	// Find the first active camera (if any)
	camZoom := 1.0
	if cam, ok := rs.Cameras.First(); ok {
		if cam.Zoom != 0 {
			camZoom = cam.Zoom
		}
	}

	// Iterate through entities that have a sprite and transform
	for e := range rs.Sprites.Entities {
		sprite, okS := rs.Sprites.Get(e)
		transform, okT := rs.Transforms.Get(e)
		if !okS || !okT || sprite.Img == nil {
			continue
		}

		opts := &ebiten.DrawImageOptions{}

		// Sprite base bounds
		bounds := sprite.Img.Bounds()
		w, h := float64(bounds.Dx()), float64(bounds.Dy())

		// Center origin before scaling
		opts.GeoM.Translate(-w/2, -h/2)

		// Combine all scales
		totalScale := sprite.Scale * transform.Scale * camZoom
		if totalScale == 0 {
			totalScale = 1.0 // fallback safety
		}

		opts.GeoM.Scale(totalScale, totalScale)

		// Apply translation
		opts.GeoM.Translate(transform.X, transform.Y)

		screen.DrawImage(sprite.Img, opts)
	}
}

// Update exists to satisfy ECS interfaces but isn’t used here.
func (rs *RenderSystem) Update(dt float64) {}

