package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var macierz1 [5][5]int
	var macierz2 [5][5]int

	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			macierz1[i][j] = rand.Intn(10)
			macierz2[i][j] = rand.Intn(10)
		}
	}

	var wynik [5][5]int

	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			wynik[i][j] = macierz1[i][j] * macierz2[i][j]
		}
	}

	fmt.Println("macierz1: ", macierz1)
	fmt.Println("macierz2: ", macierz2)
	fmt.Println("wynik: ", wynik)
}
