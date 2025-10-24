package gfx

import (
    "image"
    _ "image/png"
    "os"
    "sync"

    "github.com/hajimehoshi/ebiten/v2"
)

type AssetManager struct {
    cache map[string]*ebiten.Image
    mu    sync.Mutex
}

var Assets = &AssetManager{cache: make(map[string]*ebiten.Image)}

func (a *AssetManager) LoadImage(path string) (*ebiten.Image, error) {
    a.mu.Lock()
    defer a.mu.Unlock()
    if img, ok := a.cache[path]; ok {
        return img, nil
    }
    f, err := os.Open(path)
    if err != nil {
        return nil, err
    }
    defer f.Close()
    decoded, _, err := image.Decode(f)
    if err != nil {
        return nil, err
    }
    ebimg := ebiten.NewImageFromImage(decoded)
    a.cache[path] = ebimg
    return ebimg, nil
}

