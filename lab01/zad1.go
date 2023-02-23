// Napisz prosty program, który pyta ile masz lat,
// a następnie wypisuje liczbę miesięcy i dni które masz już za sobą.
package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Ile masz lat?")
	var age int
	fmt.Scanln(&age)
	months := age * 12
	days := int(time.Since(time.Now().AddDate(-age, 0, 0)).Hours() / 24)
	fmt.Printf("Masz już za sobą %d miesięcy lub %d dni.\n", months, days)
}
