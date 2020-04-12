package main

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"log"
	"math"
)

type Camera struct {
	X float64
	Y float64
}

const CncStyleCameraControlSpeed = 0.25

type CnCStyleCameraControl struct {
	Camera  *Camera
	StartX  int
	StartY  int
	Enabled bool
}

var cursorScrollingStartImg *ebiten.Image
var cursorDefaultImg *ebiten.Image
var cursorScrollingImg *ebiten.Image

func loadCncStyleCameraControlAsset() {
	var err error
	cursorScrollingStartImg, _, err = ebitenutil.NewImageFromFile("crosshairpack_kenney/PNG/Outline/crosshair005.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}

	cursorDefaultImg, _, err = ebitenutil.NewImageFromFile("crosshairpack_kenney/PNG/Outline/crosshair001.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}

	cursorScrollingImg, _, err = ebitenutil.NewImageFromFile("crosshairpack_kenney/PNG/Outline/crosshair022.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
}

func (e *CnCStyleCameraControl) Update() error {
	if e.stoppedControl() {
		println("stopped")
		e.Enabled = false
		return nil
	}

	if e.startedControl() {
		println("started")
		e.Enabled = true
		e.StartX, e.StartY = ebiten.CursorPosition()
	}

	if e.controlling() {
		e.moveCamera()
	}

	return nil
}

func (e *CnCStyleCameraControl) DrawSprite(screen *ebiten.Image, pos ebiten.GeoM) {
	if e.controlling() {
		drawCentered(
			cursorScrollingStartImg,
			screen,
			float64(e.StartX),
			float64(e.StartY),
			&ebiten.DrawImageOptions{},
		)
	}

	if e.controlling() {
		op := &ebiten.DrawImageOptions{}

		curX, curY := ebiten.CursorPosition()
		op.GeoM.Translate(-36, -36)
		diffX := curX - e.StartX
		diffY := curY - e.StartY

		if math.Abs(float64(diffX)) > math.Abs(float64(diffY)) {
			if diffX < 0 {
				op.GeoM.Rotate(270 * math.Pi / 180)
			}
			if diffX > 0 {
				op.GeoM.Rotate(90 * math.Pi / 180)
			}
		} else {
			if diffY > 0 {
				op.GeoM.Rotate(180 * math.Pi / 180)
			}
			if diffY < 0 {
				op.GeoM.Rotate(0 * math.Pi / 180)
			}
		}
		op.GeoM.Translate(+36, +36)
		drawCursor(cursorScrollingImg, screen, op)
	} else {
		drawCursor(cursorDefaultImg, screen, &ebiten.DrawImageOptions{})
	}
}

func (e *CnCStyleCameraControl) startedControl() bool {
	return !e.Enabled && ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft)
}

func (e *CnCStyleCameraControl) stoppedControl() bool {
	return e.Enabled && !ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft)
}

func (e *CnCStyleCameraControl) controlling() bool {
	return ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft)
}

func (e *CnCStyleCameraControl) moveCamera() {
	curX, curY := ebiten.CursorPosition()

	offX := e.StartX - curX
	offY := e.StartY - curY

	e.Camera.X -= float64(offX) * CncStyleCameraControlSpeed
	e.Camera.Y -= float64(offY) * CncStyleCameraControlSpeed
}

func drawCentered(img *ebiten.Image, target *ebiten.Image, tx, ty float64, op *ebiten.DrawImageOptions) {
	width, height := cursorScrollingStartImg.Size()
	op.GeoM.Translate(tx-float64(width/2), ty-float64(height/2))
	target.DrawImage(img, op)
}

func drawCursor(img *ebiten.Image, screen *ebiten.Image, op *ebiten.DrawImageOptions) {
	curX, curY := ebiten.CursorPosition()

	drawCentered(
		img,
		screen,
		float64(curX),
		float64(curY),
		op,
	)
}
