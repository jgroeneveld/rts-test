package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hajimehoshi/ebiten/inpututil"
)

type DebugEntity struct {
}

func (e *DebugEntity) Update() error {
	return nil
}

func (e *DebugEntity) DrawSprite(screen *ebiten.Image, pos ebiten.GeoM) {
	ebitenutil.DebugPrint(screen, fmt.Sprintf("TPS: %0.2f", ebiten.CurrentTPS()))
}

type FullscreenToggleEntity struct {
}

func (e *FullscreenToggleEntity) Update() error {
	if inpututil.IsKeyJustReleased(ebiten.KeyF) {
		ebiten.SetFullscreen(!ebiten.IsFullscreen())
	}
	return nil
}
