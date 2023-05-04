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
	layers [100][100]int
}

func (g *Game) Update() error {
	for i := range g.layers {
		for j := range g.layers[i] {
			if g.layers[i][j] == 1 {
				ni, nj := i+rand.Intn(3)-1, j+rand.Intn(3)-1
				if ni >= 0 && ni < len(g.layers) && nj >= 0 && nj < len(g.layers[0]) {
					g.layers[ni][nj] = 1
					g.layers[i][j] = 0
				}
			}
		}
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	tileXcount := len(g.layers[0])
	tileYcount := len(g.layers)
	x, _ := screen.Size()
	tileSize := x / tileXcount

	for i := 0; i < tileXcount; i++ {
		for j := 0; j < tileYcount; j++ {
			if g.layers[j][i] == 0 {
				ebitenutil.DrawRect(screen, float64(i*tileSize), float64(j*tileSize), float64(tileSize), float64(tileSize), color.White)
			}
		}
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 500, 500
}

func main() {
	antMap := [100][100]int{}
	for i := 0; i < 100; i++ {
		for j := 0; j < 100; j++ {
			if rand.Intn(100) < 10 {
				antMap[i][j] = rand.Intn(2)
			}

		}
	}

	game := &Game{
		layers: antMap}

	getSize()
	ebiten.SetWindowTitle("Your game's title")

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
