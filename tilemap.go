package main

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"image"
	"log"
	"math/rand"
	"time"
)

type TileEntity struct {
	GeoM       ebiten.GeoM
	Map        TileMap
	WaterTicks int
}

func (e *TileEntity) Update() error {
	e.AnimateWater()

	return nil
}

func (e *TileEntity) AnimateWater() {
	e.WaterTicks++

	//if e.WaterTicks > 15 {
	//	for y, col := range e.Map {
	//		for x, tileType := range col {
	//			if tileType == 18 {
	//				e.Map[y][x] = 19
	//			}
	//
	//			if tileType == 19 {
	//				e.Map[y][x] = 18
	//			}
	//		}
	//	}
	//	e.WaterTicks = 0
	//}
}

func (e *TileEntity) DrawSprite(screen *ebiten.Image, pos ebiten.GeoM)  {
	tileSize := 64

	for _, layer := range e.Map {
		for y, col := range layer {
			for x, tileType := range col {
				if tileType >= 0 {
					op := &ebiten.DrawImageOptions{GeoM: pos}
					op.GeoM.Translate(float64(x*tileSize), float64(y*tileSize))

					tileXNum := 18
					sx := (tileType % tileXNum) * tileSize
					sy := (tileType / tileXNum) * tileSize

					tileImage := tileSheet.SubImage(image.Rect(sx, sy, sx+tileSize, sy+tileSize))
					screen.DrawImage(tileImage.(*ebiten.Image), op)
				}
			}
		}
	}
}

var tileImages []*ebiten.Image
var tileSheet *ebiten.Image

type TileLayer [][]int

func (m TileLayer) Width() int {
	return len(m)
}

func (m TileLayer) Height() int {
	if len(m) == 0 {
		return 0
	}
	return len(m[0])
}

type TileMap []TileLayer

func (m TileMap) Width() int {
	if len(m) == 0 {
		return 0
	}
	return m[0].Width()
}

func (m TileMap) Height() int {
	if len(m) == 0 {
		return 0
	}
	return m[0].Height()
}

func init() {
	rand.Seed(int64(time.Now().Nanosecond()))
	//for i := 0; i < 42; i++ {
	//	imageName := fmt.Sprintf("scifiTile_%02d.png", i+1)
	//	img, _, err := ebitenutil.NewImageFromFile("kenney_rtssci-fi/PNG/Default size/Tile/" +imageName, ebiten.FilterDefault)
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//	tileImages = append(tileImages, img)
	//}
	var err error
	tileSheet, _, err = ebitenutil.NewImageFromFile("kenney_rtssci-fi/Tilesheet/scifi_tilesheet.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}

	loadCncStyleCameraControlAsset()
}

const LOGIC_FREE = 0
const LOGIC_BLOCKED = 1

func generateMap(width, height int) TileMap {
	logicLayer := emptyLayer(width, height)
	groundLayer := randomGroundTiles(width, height, logicLayer)
	landscapeLayer := randomTrees(width, height, logicLayer)
	customLayer := randomBuildings(width, height, logicLayer)

	return TileMap{
		groundLayer,
		landscapeLayer,
		customLayer,
	}
}

func emptyLayer(width int, height int) TileLayer {
	layer := make(TileLayer, height)
	for i := range layer {
		layer[i] = make([]int, width)
		for j := range layer[i] {
			layer[i][j] = -1
		}
	}
	return layer
}

func randomGroundTiles(width int, height int, logicLayer TileLayer) TileLayer {
	layer := emptyLayer(width, height)

	for y, col := range layer {
		for x, _ := range col {
			layer[y][x] = randomIntBetween(0, 1)
			logicLayer[y][x] = LOGIC_FREE
		}
	}

	return layer
}

func placeRandomTiles(count int, tileTypes []int, layer TileLayer, logicLayer TileLayer) {
	repeat(count, func() {
		x := randomIntBetween(1, layer.Width()-1)
		y := randomIntBetween(1, layer.Height()-1)

		if logicLayer[y][x] == LOGIC_FREE {
			layer[y][x] = randomSliceInt(tileTypes)
			logicLayer[y][x] = LOGIC_BLOCKED
		}
	})
}

func randomTrees(width int, height int, logicLayer TileLayer) TileLayer {
	layer := emptyLayer(width, height)

	tileTypes := []int{2, 3, 20, 21, 36, 37, 38, 39}
	placeRandomTiles(randomIntBetween(20, 100), tileTypes, layer, logicLayer)

	return layer
}

func randomBuildings(width int, height int, logicLayer TileLayer) TileLayer {
	layer := emptyLayer(width, height)

	tileTypes := []int{32, 33, 34, 35, 50, 51, 52, 53}
	placeRandomTiles(randomIntBetween(10, 30), tileTypes, layer, logicLayer)

	return layer
}

func repeat(count int, f func()) {
	for i := 0; i < count; i++ {
		f()
	}
}

func randomIntBetween(min, max int) int {
	top := max - min
	r := rand.Intn(top+1)
	return min + r
}

func randomSliceInt(numbers []int) int {
	i := randomIntBetween(0, len(numbers)-1)
	return numbers[i]
}
