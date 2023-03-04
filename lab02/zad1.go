package main

import (
	"fmt"
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

	fmt.Println("Wartości ASCII: ", ascii)
	var found bool

	var i int
	i = 1

	for found == false {
		silnia := strconv.Itoa(silnia(i))
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
			fmt.Println("Znaleziono silną liczbę: ", i)
			found = true
			break
		}

		i = i + 1
	}

}

func silnia(n int) int {
	if n == 0 {
		return 1
	}
	return n * silnia(n-1)
}

func contains(str string, substr string) bool {
	fmt.Println(str, substr)
	return len(str) > 0 && len(substr) > 0 && len(str) >= len(substr) && (str == substr || (str[0:len(substr)] == substr || contains(str[1:], substr)))
}
