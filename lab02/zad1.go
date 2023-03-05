package main

import (
	"fmt"
	"math"
	"math/big"
	"strconv"
)

func main() {
	var imie string
	var nazwisko string

	fmt.Println("Podaj imie: ")
	fmt.Scan(&imie)
	fmt.Println("Podaj nazwisko: ")
	fmt.Scan(&nazwisko)

	imie = imie[0:3]
	nazwisko = nazwisko[0:3]
	var nick string
	nick = imie + nazwisko

	fmt.Println("Twój nick to: ", nick)

	var ascii []int
	for i := 0; i < len(nick); i++ {
		ascii = append(ascii, int(nick[i]))
	}

	// fmt.Println("Wartości ASCII: ", ascii)
	var found bool

	var i int
	i = 1

	for found == false {
		silnia := silnia2(i).String()
		digits := ""
		for _, ch := range silnia {
			digits += string(ch)
		}

		allFound := true
		for _, code := range ascii {
			if !contains(digits, strconv.Itoa(code)) {
				allFound = false
				break
			}
		}

		if allFound {
			fmt.Println("Silna liczba: ", i)
			found = true
			var calls [31]int
			fibonacci(30, &calls)
			var closest int
			for j := 0; j < len(calls); j++ {
				if math.Abs(float64(i-calls[j])) < math.Abs(float64(i-calls[closest])) {
					closest = j
				}
			}
			fmt.Println("Słaba liczba: ", closest)
		}

		i = i + 1
	}

}

func silnia(n int) int {
	if n == 0 {
		return 1
	}
	fmt.Print(n)
	return n * silnia(n-1)
}

func silnia2(n int) *big.Int {
	if n == 0 {
		return big.NewInt(1)
	}
	return big.NewInt(int64(n)).Mul(big.NewInt(int64(n)), silnia2(n-1))
}

func contains(str string, substr string) bool {
	if len(str) < len(substr) {
		return false
	}
	for i := 0; i < len(str)-len(substr)+1; i++ {
		if str[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}

func fibonacci(n int, calls *[31]int) int {
	calls[n] += 1
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return fibonacci(n-1, calls) + fibonacci(n-2, calls)
}
