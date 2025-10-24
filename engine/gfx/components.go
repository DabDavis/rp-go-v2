package gfx

import "github.com/hajimehoshi/ebiten/v2"

type Transform struct {
    X, Y     float64
    Scale    float64
    Rotation float64
}

type Sprite struct {
    Img *ebiten.Image
}

type Camera struct {
    X, Y float64
    Zoom float64
}
