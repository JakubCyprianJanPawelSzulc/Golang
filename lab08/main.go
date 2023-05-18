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
	layers [100][100]Ant
}

type Ant struct {
	direction   int
	timeToMove  int
	currentMove int
}

func (g *Game) Update() error {
	for tileX := range g.layers {
		for tileY := range g.layers[tileX] {
			if g.layers[tileX][tileY].timeToMove != 0 {
				if g.layers[tileX][tileY].currentMove < g.layers[tileX][tileY].timeToMove {
					g.layers[tileX][tileY].currentMove++
				} else {
					if g.layers[tileX][tileY].direction == 0 {
						if tileY != 0 && g.layers[tileX][tileY-1].timeToMove == 0 {
							g.layers[tileX][tileY-1] = g.layers[tileX][tileY]
							g.layers[tileX][tileY] = Ant{}
							g.layers[tileX][tileY-1].currentMove = 0
						}
					}
					if g.layers[tileX][tileY].direction == 1 {
						if tileY != 99 && g.layers[tileX][tileY+1].timeToMove == 0 {
							g.layers[tileX][tileY+1] = g.layers[tileX][tileY]
							g.layers[tileX][tileY] = Ant{}
							g.layers[tileX][tileY+1].currentMove = 0
						}
					}
					if g.layers[tileX][tileY].direction == 2 {
						if tileX != 0 && g.layers[tileX-1][tileY].timeToMove == 0 {
							g.layers[tileX-1][tileY] = g.layers[tileX][tileY]
							g.layers[tileX][tileY] = Ant{}
							g.layers[tileX-1][tileY].currentMove = 0
						}
					}
					if g.layers[tileX][tileY].direction == 3 {
						if tileX != 99 && g.layers[tileX+1][tileY].timeToMove == 0 {
							g.layers[tileX+1][tileY] = g.layers[tileX][tileY]
							g.layers[tileX][tileY] = Ant{}
							g.layers[tileX+1][tileY].currentMove = 0
						}
					}
					newDirection := rand.Intn(4)
					for newDirection == g.layers[tileX][tileY].direction {
						newDirection = rand.Intn(4)
					}
					g.layers[tileX][tileY].direction = newDirection
					g.layers[tileX][tileY].currentMove = 0
				}
			}
		}
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {

	for tileX := range g.layers {
		for tileY := range g.layers[tileX] {
			if g.layers[tileX][tileY].timeToMove != 0 {
				ebitenutil.DrawRect(screen, float64(tileX*5), float64(tileY*5), 5, 5, color.White)
			}
		}
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 500, 500
}

func main() {

	antMap := [100][100]Ant{}
	for i := 0; i < 100; i++ {
		for j := 0; j < 100; j++ {
			if rand.Intn(100) < 10 {
				antMap[i][j] = Ant{rand.Intn(4), 1, 0}
			}
		}
	}

	game := &Game{
		layers: antMap}

	getSize()
	ebiten.SetWindowTitle("mrówy")

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
