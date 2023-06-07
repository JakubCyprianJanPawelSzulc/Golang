package main

import (
	"fmt"
	"math/rand"
	"time"
)

func SprawdzaniePlanszy(plansza [][]int) bool {
	r := 0
	for v := 0; v < 5; v++ {
		for t := 0; t < 5; t++ {
			if plansza[v][t] == 0 {
				r = r + 1
			}
		}
	}
	if r == 0 {
		return true
	} else {
		return false
	}
}

func RysujPlansze(plansza [][]int) {
	for i := 0; i < len(plansza); i++ {
		if i == 0 {
			fmt.Println("    0   1   2   3   4")
			fmt.Println(" -----------------------")
		}
		for j := 0; j < len(plansza[i]); j++ {
			if j == 0 {
				fmt.Print(i, " ")
			}
			fmt.Print("| ")
			if plansza[i][j] == 0 {
				fmt.Print(" ")
			} else {
				fmt.Print(plansza[i][j])
			}
			fmt.Print(" ")
			if j == 4 {
				fmt.Print("|")
			}
		}
		if i == 4 {
			fmt.Println()
			fmt.Println(" -----------------------")
		} else {
			fmt.Println()
		}
	}
}

func WyborPol(col int, row int, plansza [][]int) {
	plansza[row][col] = -1
}

func Strzelanie(plansza [][]int, pozycje [][]int, licznik *[]int, czas float64, done chan bool) {
	listapomocnicza := [][]int{}
	col := rand.Intn(5)
	row := rand.Intn(5)
	col2 := rand.Intn(5)
	row2 := rand.Intn(5)

	z := 0
	s := 0
	o := 0
	for q := 0; q < len(pozycje); q++ {
		if col == pozycje[q][0] && row == pozycje[q][1] {
			z = z + 1
		}
	}
	if z == 0 {
		listapomocnicza = append(listapomocnicza, []int{col, row})
		for y := 0; y < len(pozycje); y++ {
			if col2 == pozycje[y][0] && row2 == pozycje[y][1] {
				s = s + 1
			}
		}
		for l := 0; l < len(listapomocnicza); l++ {
			if col2 == listapomocnicza[l][0] && row2 == listapomocnicza[l][1] {
				o = o + 1
			}
		}
	}

	if z == 0 && s == 0 && o == 0 {
		WyborPol(col, row, plansza)
		pozycje = append(pozycje, []int{col, row})
		WyborPol(col2, row2, plansza)
		pozycje = append(pozycje, []int{col2, row2})
		RysujPlansze(plansza)

		timeout := time.Duration(czas) * time.Second
		timer := time.NewTimer(timeout)
		defer timer.Stop()

		go func() {
			var a string
			fmt.Scanln(&a)
			dlugosc := len(a)
			if dlugosc == 2 {
				x := int(a[0] - '0')
				y := int(a[1] - '0')
				hit := false
				for k := 0; k < len(pozycje); k++ {
					if x == pozycje[k][0] && y == pozycje[k][1] {
						fmt.Println("hit")
						pozycje = append(pozycje[:k], pozycje[k+1:]...)
						plansza[y][x] = 0
						RysujPlansze(plansza)
						*licznik = append(*licznik, 1)
						hit = true
						break
					}
				}
				if !hit {
					fmt.Println("miss")
				}
			} else {
				fmt.Println("error")
			}
			done <- true
		}()

		select {
		case <-done:
			return
		case <-timer.C:
			return
		}
	}
}

func Gra() {
	var czas float64
	fmt.Print("Ile sekund na turę? ")
	fmt.Scanln(&czas)
	var trybgry int = 2
	// fmt.Print("2 czy 3 przeciwników na turę? ")
	// fmt.Scanln(&trybgry)
	licznik := []int{}
	pozycje := [][]int{}
	plansza := [][]int{{0, 0, 0, 0, 0}, {0, 0, 0, 0, 0}, {0, 0, 0, 0, 0}, {0, 0, 0, 0, 0}, {0, 0, 0, 0, 0}}
	if trybgry == 2 {
		for !SprawdzaniePlanszy(plansza) {
			rr := 0
			for dd := 0; dd < len(plansza); dd++ {
				for tt := 0; tt < 5; tt++ {
					if plansza[dd][tt] == 0 {
						rr = rr + 1
					}
				}
			}
			if rr == 1 {
				for h := 0; h < len(plansza); h++ {
					for c := 0; c < 5; c++ {
						if plansza[h][c] == 0 {
							plansza[h][c] = -1
							pozycje = append(pozycje, []int{h, c})
						}
					}
				}
			}
			Strzelanie(plansza, pozycje, &licznik, czas, make(chan bool))
		}
	}
	RysujPlansze(plansza)
	wynik := len(licznik)
	fmt.Println("twoj wynik: ", wynik)
	fmt.Print("Chcesz zagrać ponownie? (T/N) ")
	var odpowiedz string
	fmt.Scanln(&odpowiedz)
	if odpowiedz == "T" || odpowiedz == "t" || odpowiedz == "Tak" || odpowiedz == "tak" {
		Gra()
	} else {
		fmt.Println("Dziękuję za zagranie w moją grę")
	}
}

func Instrukcja() {
	planszademo := [][]int{{0, 0, 0, 0, 0}, {0, 0, -1, 0, 0}, {0, 0, 0, 0, 0}, {0, 0, 0, 0, -1}, {0, 0, 0, 0, 0}}
	planszademo2 := [][]int{{0, 0, 0, 0, 0}, {0, 0, 0, 0, 0}, {0, 0, 0, 0, 0}, {0, 0, 0, 0, -1}, {0, 0, 0, 0, 0}}
	planszademo3 := [][]int{{0, 0, 0, 0, 0}, {0, 0, 0, 0, 0}, {0, 0, 0, 0, 0}, {0, 0, 0, 0, 0}, {0, 0, 0, 0, 0}}
	fmt.Println("Działanie gry:")
	RysujPlansze(planszademo)
	time.Sleep(1 * time.Second)
	fmt.Println("Wpisujemy współrzędne w następny sposób:")
	time.Sleep(2 * time.Second)
	fmt.Println("21")
	fmt.Println("Przeciwnik o współrzędnych 2 1 zostaje zmieciony z planszy")
	time.Sleep(1 * time.Second)
	RysujPlansze(planszademo2)
	time.Sleep(1 * time.Second)
	fmt.Println("43")
	time.Sleep(1 * time.Second)
	RysujPlansze(planszademo3)
	time.Sleep(3 * time.Second)
}

func Zasadygry() {
	fmt.Println("Zasady gry: Nowi przeciwnicy pojawiają się co każdą turę")
	fmt.Println("Twoim zadaniem jest strzelanie do nich poprzez wpisywanie ich współrzędnych")
	fmt.Println("Gra kończy się gdy przeciwnicy opanują całą planszę")
	fmt.Println("*****************************************************************************")
	time.Sleep(3 * time.Second)
	fmt.Print("Czy chcesz obejrzeć dokładniejszą instrukcję? (T/N) ")
	var czyinstrukcja string
	fmt.Scanln(&czyinstrukcja)
	if czyinstrukcja == "T" || czyinstrukcja == "t" || czyinstrukcja == "Tak" || czyinstrukcja == "tak" {
		Instrukcja()
		Gra()
	} else {
		Gra()
	}
}

func main() {
	Zasadygry()
}
