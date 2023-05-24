package main

import (
	"bufio"
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
	if len(os.Args) < 2 {
		fmt.Println("Podaj nazwę pliku jako argument wywołania programu.")
		return
	}
	fileName := os.Args[1]

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
		color := getColor(status)

		fmt.Printf("%s%s%s\n", color, line, Reset)
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Błąd odczytu pliku: %v\n", err)
	}
}

func getStatus(line string) string {
	if strings.HasPrefix(line, "[x]") {
		return "DONE"
	} else if strings.HasPrefix(line, "[-]") {
		return "NOPE"
	} else if strings.HasPrefix(line, "[+]") {
		return "IN PROGRESS"
	} else if strings.HasPrefix(line, "[ ]") {
		return "TODO"
	} else {
		return "COMMENT"
	}
}

func getColor(status string) string {
	switch status {
	case "DONE":
		return Green
	case "NOPE":
		return Gray
	case "IN PROGRESS":
		return Red
	case "TODO":
		return Yellow
	default:
		return Reset
	}
}
