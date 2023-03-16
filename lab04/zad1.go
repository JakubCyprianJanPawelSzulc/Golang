package main

import (
	"fmt"
)

func main() {
	var array1 [20]float64
	var array2 [20]float64

	var suma float64 = 0.0
	var suma1 float64 = 0.0
	var suma2 float64 = 0.0

	for i := 0; i < 20; i++ {
		array1[i] = 2.0
		array2[i] = 3.0
	}

	for i := 0; i < 20; i++ {
		suma += array1[i] + array2[i]
		suma1 += array1[i]
		suma2 += array2[i]
	}

	fmt.Println(suma)
	fmt.Println(suma1)
	fmt.Println(suma2)
}
