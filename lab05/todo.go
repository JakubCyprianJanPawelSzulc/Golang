package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

const (
	Reset  = "\033[0m"
	Gray   = "\033[1;30m"
	Yellow = "\033[1;33m"
	Green  = "\033[1;32m"
	Red    = "\033[1;31m"
)

func main() {
	noColor := flag.Bool("nocolor", false, "Wyłącza kolorowanie wyjścia")
	showStatus := flag.String("status", "", "Wyświetla tylko zadania o określonym statusie (done, nope, inprogress, todo)")
	flag.Parse()

	if flag.NArg() < 1 {
		fmt.Println("Podaj nazwę pliku jako argument wywołania programu.")
		return
	}

	fileName := flag.Arg(0)

	file, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("Błąd podczas otwierania pliku: %v\n", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		status := getStatus(line)
		if *showStatus != "" && status != *showStatus {
			continue
		}
		if *noColor {
			fmt.Println(line)
		} else {
			color := getColor(status)
			fmt.Printf("%s%s%s\n", color, line, Reset)
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Printf("Błąd odczytu pliku: %v\n", err)
	}
}

func getStatus(line string) string {
	if strings.HasPrefix(line, "[x]") {
		return "done"
	} else if strings.HasPrefix(line, "[-]") {
		return "nope"
	} else if strings.HasPrefix(line, "[+]") {
		return "inprogress"
	} else if strings.HasPrefix(line, "[ ]") {
		return "todo"
	} else {
		return "comment"
	}
}

func getColor(status string) string {
	switch status {
	case "done":
		return Green
	case "nope":
		return Gray
	case "inprogress":
		return Red
	case "todo":
		return Yellow
	default:
		return Reset
	}
}
