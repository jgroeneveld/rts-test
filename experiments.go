package main

import (
	"github.com/hajimehoshi/ebiten"
	"log"
)

type ScreenInverterEntity struct {
	showInverted bool
	count        int
}

func (e *ScreenInverterEntity) Update() error {
	e.count++

	if e.count > 5 {
		e.showInverted = !e.showInverted
		e.count = 0
	}

	return nil
}

func (e *ScreenInverterEntity) DrawSprite(screen *ebiten.Image, pos ebiten.GeoM) {
	if !e.showInverted {
		return
	}

	op := &ebiten.DrawImageOptions{}
	op.ColorM.Scale(-1, -1, -1, 1)
	op.ColorM.Translate(1, 1, 1, 0)
	sw, sh := screen.Size()

	temp, err := ebiten.NewImage(sw, sh, ebiten.FilterDefault)

	if err != nil {
		log.Fatal(err)
	}

	temp.DrawImage(screen, op)
	screen.DrawImage(temp, &ebiten.DrawImageOptions{})
}