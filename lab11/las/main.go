package main

import (
	"bytes"
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/hajimehoshi/go-mp3"
	"github.com/hajimehoshi/oto/v2"
)

type Position struct {
	X, Y int
}

func main() {
	rand.Seed(time.Now().UnixNano())

	width := 10
	height := 10
	perc := 60

	forest := make([][]int, height)
	for i := range forest {
		forest[i] = make([]int, width)
	}

	generateForest(&forest, width, height, perc)

	fmt.Println("Drücken Sie die Eingabetaste, um einen Blitz aufzunehmen")

	for {
		var input string
		fmt.Scanln(&input)
		if input == "" {
			roundOfBurningForest(&forest, width, height)
		} else {
			break
		}
	}

	printForest(forest)
}

func printForest(forest [][]int) {
	for _, row := range forest {
		for _, cell := range row {
			if cell == 1 {
				fmt.Print("🌲")
			} else if cell == 2 {
				fmt.Print("⬛")
			} else if cell == 3 {
				fmt.Print("🔥")
			} else if cell == 4 {
				fmt.Print("💥")
			} else if cell == 5 {
				fmt.Print("⚡️")
			}
		}
		fmt.Println()
	}
	fmt.Println("")
}

func burnForest(forest *[][]int, x, y int, change *bool) {

	if x < 0 || x >= len((*forest)[0]) || y < 0 || y >= len(*forest) {
		return
	}

	if (*forest)[y][x] == 1 || (*forest)[y][x] == 4 {
		(*forest)[y][x] = 3
		*change = true
	} else {
		return
	}

	burnForest(forest, x-1, y, change)
	burnForest(forest, x+1, y, change)
	burnForest(forest, x, y-1, change)
	burnForest(forest, x, y+1, change)
}

func clearFire(forest *[][]int) {
	for y, row := range *forest {
		for x := range row {
			if (*forest)[y][x] == 3 || (*forest)[y][x] == 4 || (*forest)[y][x] == 5 {
				(*forest)[y][x] = 2
			}
		}
	}
}

func generateForest(forest *[][]int, width, height, perc int) {
	treeCount := (width * height * perc) / 100
	count := 0

	for count < treeCount {
		x := rand.Intn(width)
		y := rand.Intn(height)

		if (*forest)[y][x] == 0 {
			(*forest)[y][x] = 1
			count++
		}
	}

	for y, row := range *forest {
		for x := range row {
			if row[x] != 1 {
				(*forest)[y][x] = 2
			}
		}
	}
}

func roundOfBurningForest(forest *[][]int, width, height int) {
	lightningX := rand.Intn(width)
	lightningY := rand.Intn(height)
	change := false

	if (*forest)[lightningY][lightningX] == 1 {
		(*forest)[lightningY][lightningX] = 4
	} else if (*forest)[lightningY][lightningX] == 2 {
		(*forest)[lightningY][lightningX] = 5
	}

	go playThunderSound()

	printForest(*forest)
	burnForest(forest, lightningX, lightningY, &change)
	if change {
		printForest(*forest)
		fmt.Println("Spłonęło", countBurnedTrees(*forest))
	}
	clearFire(forest)
}

func countBurnedTrees(forest [][]int) int {
	count := 0
	for _, row := range forest {
		for _, cell := range row {
			if cell == 4 || cell == 3 {
				count++
			}
		}
	}
	return count
}

func playThunderSound() {
	// fileBytes, err := os.ReadFile("./thunder1.mp3")
	fileBytes, err := os.ReadFile("./thunder2.mp3")
	if err != nil {
		panic("Das Öffnen der Datei ist fehlgeschlagen: " + err.Error())
	}
	fileBytesReader := bytes.NewReader(fileBytes)
	decodedMp3, err := mp3.NewDecoder(fileBytesReader)
	if err != nil {
		panic("mp3.NewDecoder ist kaputt: " + err.Error())
	}

	samplingRate := 44100
	numOfChannels := 2
	audioBitDepth := 2

	otoCtx, readyChan, err := oto.NewContext(samplingRate, numOfChannels, audioBitDepth)
	if err != nil {
		panic("oto.NewContext funktioniert nicht: " + err.Error())
	}
	<-readyChan

	player := otoCtx.NewPlayer(decodedMp3)

	player.Play()

	for player.IsPlaying() {
		time.Sleep(time.Millisecond)
	}
	err = player.Close()
	if err != nil {
		panic("player.Close ist fehlgeschlagen: " + err.Error())
	}
}
