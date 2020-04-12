package main

import (
	"github.com/hajimehoshi/ebiten"
)

type EntityManager struct {
	Camera *Camera
	Entities []Entity
}

func (m *EntityManager) Update() error {
	for _, entity := range m.Entities {
		updatable, ok := entity.(Updater)
		if ok {
			err := updatable.Update()
			if err != nil {
				return err
			}
		}
	}

	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		m.Camera.X -= 10
	}

	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		m.Camera.X += 10
	}

	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		m.Camera.Y -= 10
	}

	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		m.Camera.Y += 10
	}

	return nil
}

func (m *EntityManager) Draw(screen *ebiten.Image) {
	m.drawWorld(screen)
	m.drawUI(screen)
}


func (m *EntityManager) drawWorld(screen *ebiten.Image) {
	pos := ebiten.GeoM{}
	pos.Translate(-m.Camera.X, -m.Camera.Y)

	for _, entity := range m.Entities {
		sprite, ok := entity.(Sprite)
		if ok {
			sprite.DrawSprite(screen, pos)
		}
	}
}

func (m *EntityManager) drawUI(screen *ebiten.Image) {
	for _, entity := range m.Entities {
		ui, ok := entity.(UIElement)
		if ok {
			ui.DrawUI(screen)
		}
	}
}
type Entity interface {
}

type Updater interface {
	Update() error
}

type Sprite interface {
	DrawSprite(screen *ebiten.Image, pos ebiten.GeoM)
}

type UIElement interface {
	DrawUI(screen *ebiten.Image)
}
