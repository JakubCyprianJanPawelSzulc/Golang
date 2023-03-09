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
	var tab []int
	var tab2 []int
	for i := start; i <= end-1; i += step {
		var maxLength, maxNum int
		for j := i; j <= i+step; j++ {
			length := collatz(j)
			if length > maxLength {
				maxLength = length
				maxNum = j
			}
		}
		tab = append(tab, maxLength)
		tab2 = append(tab2, maxNum)
	}
	fmt.Println(tab)
	fmt.Println(tab2)
	//srednia z wynikow
	var sum int
	for _, v := range tab {
		sum += v
	}
	fmt.Println("Srednia: ", float64(sum)/float64(len(tab)))
	//mediana z wynikow
	var median float64
	if len(tab)%2 == 0 {
		median = float64(tab[len(tab)/2]+tab[len(tab)/2-1]) / 2
	} else {
		median = float64(tab[len(tab)/2])
	}
	fmt.Println("Mediana: ", median)
	//odchylenie standardowe wynikow
	var sum2 float64
	for _, v := range tab {
		sum2 += float64(v) * float64(v)
	}
	fmt.Println("Odchylenie standardowe: ", sum2/float64(len(tab))-float64(sum*sum)/float64(len(tab)*len(tab)))
	//wariancja wynikow
	fmt.Println("Wariancja: ", sum2/float64(len(tab))-float64(sum*sum)/float64(len(tab)*len(tab)))
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
