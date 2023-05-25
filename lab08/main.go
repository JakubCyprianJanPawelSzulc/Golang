package main

import (
	"flag"
	"image/color"
	"log"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct {
	layers    [100][100]Tile
	leafLayer [100][100]bool
}

type Tile struct {
	ant Ant
}

type Ant struct {
	direction   int
	timeToMove  int
	currentMove int
}

func (g *Game) Update() error {
	for tileX := range g.layers {
		for tileY := range g.layers[tileX] {
			if g.layers[tileX][tileY].ant.timeToMove != 0 {
				if g.layers[tileX][tileY].ant.currentMove < g.layers[tileX][tileY].ant.timeToMove {
					g.layers[tileX][tileY].ant.currentMove++
				} else {
					ant := g.layers[tileX][tileY].ant
					nextTileX, nextTileY := getNextTileCoordinates(tileX, tileY, ant.direction)
					if isValidTile(nextTileX, nextTileY) && g.layers[nextTileX][nextTileY].ant.timeToMove == 0 {
						if g.leafLayer[nextTileX][nextTileY] {
							g.leafLayer[nextTileX][nextTileY] = false
						}

						g.layers[nextTileX][nextTileY].ant = ant
						g.layers[tileX][tileY].ant = Ant{}
						g.layers[nextTileX][nextTileY].ant.currentMove = 0
					} else {
						newDirection := rand.Intn(4)
						for newDirection == ant.direction {
							newDirection = rand.Intn(4)
						}
						g.layers[tileX][tileY].ant.direction = newDirection
						g.layers[tileX][tileY].ant.currentMove = 0
					}
				}
			}
		}
	}

	return nil
}

func getNextTileCoordinates(tileX, tileY, direction int) (int, int) {
	switch direction {
	case 0:
		return tileX, tileY - 1
	case 1:
		return tileX, tileY + 1
	case 2:
		return tileX - 1, tileY
	case 3:
		return tileX + 1, tileY
	}
	return tileX, tileY
}

func isValidTile(tileX, tileY int) bool {
	return tileX >= 0 && tileX < 100 && tileY >= 0 && tileY < 100
}

func (g *Game) Draw(screen *ebiten.Image) {
	for tileX := range g.layers {
		for tileY := range g.layers[tileX] {
			if g.layers[tileX][tileY].ant.timeToMove != 0 {
				ebitenutil.DrawRect(screen, float64(tileX*5), float64(tileY*5), 5, 5, color.Black)
			} else if g.leafLayer[tileX][tileY] == true {
				ebitenutil.DrawRect(screen, float64(tileX*5), float64(tileY*5), 5, 5, color.RGBA{0, 255, 0, 255})
			} else if g.layers[tileX][tileY].ant.timeToMove == 0 {
				ebitenutil.DrawRect(screen, float64(tileX*5), float64(tileY*5), 5, 5, color.White)
			}
		}
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 500, 500
}

func main() {
	antMap := [100][100]Tile{}
	leafLayer := [100][100]bool{}
	for i := 0; i < 100; i++ {
		for j := 0; j < 100; j++ {
			if rand.Intn(100) < 10 {
				antMap[i][j].ant = Ant{rand.Intn(4), 1, 0}
			}
			if rand.Intn(100) < 5 {
				leafLayer[i][j] = true
			}
		}
	}

	game := &Game{
		layers:    antMap,
		leafLayer: leafLayer,
	}

	getSize()
	ebiten.SetWindowTitle("Mrówy")

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}

func getSize() {
	w := flag.Int("w", 640, "window width")
	h := flag.Int("h", 480, "window height")
	flag.Parse()
	ebiten.SetWindowSize(*w, *h)
}
