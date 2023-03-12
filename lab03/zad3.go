package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
)

func main() {
	var start, end, step int
	start = 1
	end = 100000
	step = 100

	file, err := os.Create("collatz.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	for i := start; i <= end; i += step {
		for j := i; j <= i+step-1; j++ {
			length := collatz(j)
			file.WriteString(strconv.Itoa(j) + " " + strconv.Itoa(length) + "\n")
		}
	}

	cmd := exec.Command("gnuplot", "-p", "-e", "plot 'collatz.txt' with points")
	if err := cmd.Run(); err != nil {
		fmt.Println(err)
		return
	}
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
