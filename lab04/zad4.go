package main

import "fmt"

func main() {
	wycinek := [][]int{{4, 2, 0}, {2, 1, 1}, {5, 6, 9}}

	odwrotnosc := make([][]int, 3)
	for i := range odwrotnosc {
		odwrotnosc[i] = make([]int, 3)
	}
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			odwrotnosc[i][j] = wycinek[2-i][2-j]
		}
	}

	suma := make([][]int, 3)
	for i := range suma {
		suma[i] = make([]int, 3)
	}
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			suma[i][j] = wycinek[i][j] + odwrotnosc[i][j]
		}
	}

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf("%d ", suma[i][j])
		}
		fmt.Printf("\n")
	}
}
