package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Position struct {
	X, Y int
}

func main() {
	rand.Seed(time.Now().UnixNano())

	width := 20
	height := 20

	forest := make([][]bool, height)
	for i := range forest {
		forest[i] = make([]bool, width)
	}

	lightningX := rand.Intn(width)
	lightningY := rand.Intn(height)

	forest[lightningY][lightningX] = true

	burnForest(&forest, lightningX, lightningY)

	printForest(&forest)
}

func printForest(forest *[][]bool) {
	for _, row := range *forest {
		for _, cell := range row {
			if cell == true {
				// fmt.Print("🔥")
				fmt.Print("⬛")
			} else {
				fmt.Print("🌲")
			}
		}
		fmt.Println()
	}
}

func burnForest(forest *[][]bool, x, y int) {
	if x < 0 || x >= len((*forest)[0]) || y < 0 || y >= len(*forest) || !(*forest)[y][x] {
		return
	}
	(*forest)[y][x] = false

	burnForest(forest, x-1, y)
	burnForest(forest, x+1, y)
	burnForest(forest, x, y-1)
	burnForest(forest, x, y+1)
}
