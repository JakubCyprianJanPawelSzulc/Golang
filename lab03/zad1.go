package main

import "fmt"

func main() {

	var start, end, step int
	fmt.Print("Podaj początek zakresu: ")
	fmt.Scan(&start)
	fmt.Print("Podaj koniec zakresu: ")
	fmt.Scan(&end)
	fmt.Println("Podaj krok: ")
	fmt.Scan(&step)

	for i := start; i <= end-1; i += step {
		var maxLength, maxNum int
		for j := i; j <= i+step; j++ {
			length := collatz(j)
			if length > maxLength {
				maxLength = length
				maxNum = j
			}
		}
		fmt.Printf("Dla zakresu %d-%d: %d , długość %d.\n", i, i+step, maxNum, maxLength)
	}
}

func collatz(n int) int {
	length := 1
	for n != 1 {
		if n%2 == 0 {
			n /= 2
		} else {
			n = n*3 + 1
		}
		length++
	}
	return length
}
