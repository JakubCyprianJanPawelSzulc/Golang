// Napisz program, który wykona obliczenie pierwiastków trójmianu kwadratowego.
// Dla przypadku gdy delta jest ujemna program może wypisać komunikat, że nie ma pierwiastków
// (lub w trudniejszej wersji - może policzyć pierwiastki zespolone).

package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Podaj współczynniki trójmianu kwadratowego:")
	var a, b, c float64
	fmt.Scanln(&a, &b, &c)
	delta := b*b - 4*a*c
	if delta < 0 {
		fmt.Println("Brak pierwiastków")
	} else if delta == 0 {
		x := -b / (2 * a)
		fmt.Printf("Jeden pierwiastek: %f", x)
	} else {
		x1 := ((-b - math.Sqrt(delta)) / (2 * a))
		x2 := ((-b + math.Sqrt(delta)) / (2 * a))
		fmt.Printf("Dwa pierwiastki: %f, %f", x1, x2)
	}
}
