package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("collatz.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	counts := make(map[int]int)
	for i := 0; i < 10000; i++ {
		var num, length int
		fmt.Fscan(file, &num, &length)
		counts[length]++
	}

	var keys []int
	for k := range counts {
		keys = append(keys, k)
	}

	for i := 0; i < len(keys); i++ {
		for j := i + 1; j < len(keys); j++ {
			if counts[keys[i]] < counts[keys[j]] {
				keys[i], keys[j] = keys[j], keys[i]
			}
		}
	}

	for _, k := range keys {
		fmt.Printf("%d: %d\n", k, counts[k])
	}

}
