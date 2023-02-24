//Gra w Zgadywanie
//Trudność: zależnie od poziomu zadania.
//Napisz program, którego zadaniem jest przeprowadzić grę z człowiekiem, w zgadywanie wylosowanej przez komputer liczby. Zadanie można zrobić prosto, lub w sposób bardzo rozbudowany. Warianty, a właściwie "poziomy" przedstawione są poniżej.
//Poziom 1.
//Program na przedstawić użytkownikowi zadanie, np. "Teraz będziesz zgadywać liczbę, którą wylosowałem" - a następnie zapytać użytkownika o tę liczbę, np. pisząc "Podaj liczbę: ", wczytać ją, i sprawdzić, czy wylosowana wcześniej liczba jest taka sama. Jeżeli tak, mają się pojawić gratulacje, a program się kończy. Jeżeli nie, program powinien napisać czy liczba podana przez użytkownika jest "Za mała", lub "Za duża". Wtedy powtarzane są pytania aż w końcu użytkownik zgadnie lub przerwie program.
//Poziom 2.
//Napisz program jak wyżej, ale zmodyfikuj go w taki sposób, aby możliwe było podanie odpowiedzi "koniec", po której pytania o liczbę są przerywane a program natychmiast się kończy pisząc "żegnaj". Program powinien w powitaniu napisać informację, że po wpisaniu "koniec" nastąpi zakończenie działania. Poza tym - wszystko może pozostać takie, jak na poziomie pierwszym.

package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func main() {
	fmt.Println("Podaj maksimum zakresu")
	var max int
	fmt.Scanln(&max)
	rand.Seed(time.Now().UnixNano())
	number := rand.Intn(max)
	fmt.Println("Zgadnij liczbę z zakresu 0 -", max, "lub wpisz 'koniec' aby zakończyć grę")

	var guess string
	for {
		fmt.Scanln(&guess)
		if strings.ToLower(guess) == "koniec" {
			fmt.Println("Żegnaj!")
			break
		}

		guessNum, err := strconv.Atoi(guess)
		if err != nil {
			fmt.Println("To nie jest liczba, spróbuj ponownie lub wpisz 'koniec'")
			continue
		}

		if guessNum == number {
			fmt.Println("Zgadłeś!")
			break
		} else if guessNum > number {
			fmt.Println("Za dużo")
		} else {
			fmt.Println("Za mało")
		}
	}
}
