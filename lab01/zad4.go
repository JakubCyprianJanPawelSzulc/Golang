// Napisz program, który zagra w zgadywanie z użytkownikiem.
// Należy wylosować liczbę którą następnie użytkownik ma odgadnąć.
// Po wprowadzeniu liczby program ma wypisać informację czy liczba jest prawidłowa.
// Jeżeli nie, to czy jest większa czy mniejsza od wylosowanej.
// Gra kończy się gdy użytkownik zgadnie liczbę lub przerwie program.

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Podaj maksimum zakresu")
	var max int
	fmt.Scanln(&max)
	rand.Seed(time.Now().UnixNano())
	number := rand.Intn(max)
	fmt.Printf("Zgadnij liczbę z zakresu 0-%d\n", max)
	var guess int
	for {
		fmt.Scanln(&guess)
		if guess == number {
			fmt.Println("Zgadłeś!")
			break
		} else if guess > number {
			fmt.Println("Za dużo")
		} else {
			fmt.Println("Za mało")
		}
	}
}
