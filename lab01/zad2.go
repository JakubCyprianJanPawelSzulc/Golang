// Napisz program, który liczy ile lat miałby użytkownik, gdyby mieszkał na Marsie, na Wenus lub na innych planetach.
// Ile Ty możesz mieć aktualnie lat na takich planetach?
// Wiek mieszkańca danej planety to liczba pełnych obrotów planety wokół słońca.
// Sprawdź w internecie...
package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Ile masz lat?")
	var age int
	fmt.Scanln(&age)
	ageOnMars := int(time.Since(time.Now().AddDate(-age, 0, 0)).Hours() / 24 / 687)
	fmt.Printf("Na marsie miałbyś %d lat\n", ageOnMars)
	ageOnVenus := int(time.Since(time.Now().AddDate(-age, 0, 0)).Hours() / 24 / 225)
	fmt.Printf("Na wenus miałbyś %d lat\n", ageOnVenus)
}
