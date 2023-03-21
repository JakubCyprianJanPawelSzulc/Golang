package main

import "fmt"

func matrixMultiply(A [][]int, B [][]int) [][]int {
	n := len(A)
	C := make([][]int, n)
	for i := range C {
		C[i] = make([]int, n)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			for k := 0; k < n; k++ {
				C[i][j] += A[i][k] * B[k][j]
			}
		}
	}
	return C
}

func main() {
	A := [][]int{{-1, -2, 3}, {0, 2, -1}, {-1, 3, 0}}
	B := [][]int{{1, 5, 1}, {2, 1, 2}, {3, 2, 3}}

	C := matrixMultiply(A, B)
	for i := range C {
		fmt.Println(C[i])
	}
}
