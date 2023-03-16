package main

import "fmt"

func main() {
	var slice1 = make([][]int, 3)
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30}

	for i := 0; i < 3; i++ {
		slice1[i] = numbers[i*3 : (i+1)*3]
	}

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf("%d ", slice1[i][j])
		}
		fmt.Println()
	}

}
