// Gra w Zgadywanie
// Trudność: zależnie od poziomu zadania.
// Napisz program, którego zadaniem jest przeprowadzić grę z człowiekiem, w zgadywanie wylosowanej przez komputer liczby. Zadanie można zrobić prosto, lub w sposób bardzo rozbudowany. Warianty, a właściwie "poziomy" przedstawione są poniżej.
// Poziom 1.
// Program na przedstawić użytkownikowi zadanie, np. "Teraz będziesz zgadywać liczbę, którą wylosowałem" - a następnie zapytać użytkownika o tę liczbę, np. pisząc "Podaj liczbę: ", wczytać ją, i sprawdzić, czy wylosowana wcześniej liczba jest taka sama. Jeżeli tak, mają się pojawić gratulacje, a program się kończy. Jeżeli nie, program powinien napisać czy liczba podana przez użytkownika jest "Za mała", lub "Za duża". Wtedy powtarzane są pytania aż w końcu użytkownik zgadnie lub przerwie program.
// Poziom 2.
// Napisz program jak wyżej, ale zmodyfikuj go w taki sposób, aby możliwe było podanie odpowiedzi "koniec", po której pytania o liczbę są przerywane a program natychmiast się kończy pisząc "żegnaj". Program powinien w powitaniu napisać informację, że po wpisaniu "koniec" nastąpi zakończenie działania. Poza tym - wszystko może pozostać takie, jak na poziomie pierwszym.
// Poziom 3.
// Napisz program taki jak na poziomie 2, ale dodaj do niego dodatkowe pytanie na końcu, po odgadnięciu liczby, które brzmi "Czy gramy jeszcze raz? [T/N]". Gdy użytkownik wybierze odpowiedź (domyślnie tak lub nie), program rozpoczyna jeszcze raz grę z INNĄ wylosowaną liczbą, albo się kończy.
// Poziom 4.
// Program z poziomu 3 umożliwia wiele gier. Zmodyfikuj go tak, aby umożliwiał zapamiętanie w ilu próbach użytkownik odgadł liczbę. Po zakończeniu gry zapytaj użytkownika o imię, i zapisz je do struktury danych. Gdyby gra była kontynuowana kolejny raz, ponawiaj takie pytania, aby zebrać dane o wynikach różnych użytkowników. Na zakończenie programu wypisz podsumowanie: kto w ilu próbach odgadywał liczby, posortowane rosnąco według liczby prób.
// Poziom 5.
// Niech program z poziomu 4 zapisuje swoje wyniki do pliku, tworząc w ten sposób listę "Hall of fame", czyli tabelę najlepszych rezultatów. W tym pliku zapisuj liczbę prób, nazwę użytkownika oraz datę gry; możesz także zapisać zgadywaną liczbę. Taki plik należy wczytać po uruchomieniu programu i załadować do niego wczytane poprzednie rekordy. Program może wtedy po każdej grze informować użytkownika o pobiciu rekordu globalnego, lub osobistego. Pozostałe elementy mogą pozostać bez zmian. Użyj w tym programie struktur lub map, zależnie od tego, które rozwiązanie będzie bardziej pasowało do algorytmu, jaki opracujesz.

package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

type playerResult struct {
	name    string
	guesses int
	date    time.Time
	number  int
}

func main() {
	rand.Seed(time.Now().UnixNano())

	var results []playerResult

	loadResultsFromFile("data.txt", &results)

	for {
		fmt.Println("Podaj maksimum zakresu")
		var max int
		fmt.Scanln(&max)

		number := rand.Intn(max)
		fmt.Println("Zgadnij liczbę z zakresu 0 -", max, "lub wpisz 'koniec' aby zakończyć grę")

		var guess string
		var guesses int
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

			guesses++
			if guessNum == number {
				fmt.Println("Zgadłeś!")
				break
			} else if guessNum > number {
				fmt.Println("Za dużo")
			} else {
				fmt.Println("Za mało")
			}
		}

		if guesses > 0 {
			var name string
			fmt.Println("Podaj swoje imię")
			fmt.Scanln(&name)

			result := playerResult{name, guesses, time.Now(), number}
			results = append(results, result)

			saveResultToFile("data.txt", result)

			sort.Slice(results, func(i, j int) bool {
				return results[i].guesses < results[j].guesses
			})

			fmt.Println("Podsumowanie wyników:")
			for _, result := range results {
				fmt.Printf("%s odgadł(a) w %d próbach %s dla liczby %d\n", result.name, result.guesses, result.date.Format("2006-01-02"), result.number)
			}
		}

		var playAgain string
		fmt.Println("Chcesz zagrać jeszcze raz? [T/N]")
		fmt.Scanln(&playAgain)
		if strings.ToLower(playAgain) != "t" {
			fmt.Println("Żegnaj!")
			break
		}
	}
}

func loadResultsFromFile(filename string, results *[]playerResult) {
	file, err := os.Open(filename)
	if err != nil {
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Split(line, ";")
		if len(fields) == 4 {
			guesses, err := strconv.Atoi(fields[1])
			if err == nil {
				date, err := time.Parse("2006-01-02", fields[2])
				if err == nil {
					number, err := strconv.Atoi(fields[3])
					if err == nil {
						result := playerResult{fields[0], guesses, date, number}
						*results = append(*results, result)
					}
				}
			}
		}
	}
}

func saveResultToFile(filename string, result playerResult) {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return
	}
	defer file.Close()

	_, err = file.WriteString(fmt.Sprintf("%s;%d;%s;%d\n", result.name, result.guesses, result.date.Format("2006-01-02"), result.number))
}
