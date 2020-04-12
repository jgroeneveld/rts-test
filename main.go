package main

import (
	"github.com/hajimehoshi/ebiten"
	"log"
)

func NewGame() *Game {
	tileMap := generateMap(24, 24)
	camera := &Camera{}

	entityManager := &EntityManager{
		Camera: camera,
		Entities: []Entity{
			&TileEntity{Map: tileMap},
			//&ScreenInverterEntity{},
			&DebugEntity{},
			&CnCStyleCameraControl{Camera: camera},
			&FullscreenToggleEntity{},
		},
	}

	return &Game{
		EntityManager: entityManager,
	}
}

type Game struct {
	EntityManager *EntityManager
}

func (g *Game) Update(screen *ebiten.Image) error {
	err := g.EntityManager.Update()
	if err != nil {
		return err
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.EntityManager.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 1280, 960
}

func main() {
	ebiten.SetCursorVisible(false)
	ebiten.SetWindowSize(640, 480)

	game := NewGame()
	err := ebiten.RunGame(game)
	//err := ebiten.Run(update(), 1280, 960, 1, "Hello World")
	if err != nil {
		log.Fatal(err)
	}
}
