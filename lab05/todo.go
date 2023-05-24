package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
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

	fileNames := flag.Args()

	if len(fileNames) == 0 {
		err := processInput(os.Stdin, *noColor, *showStatus)
		if err != nil {
			fmt.Printf("Błąd przetwarzania danych wejściowych: %v\n", err)
		}
		return
	}

	for _, fileName := range fileNames {
		err := processFile(fileName, *noColor, *showStatus)
		if err != nil {
			fmt.Printf("Błąd przetwarzania pliku %s: %v\n", fileName, err)
		}
	}
}

func processFile(fileName string, noColor bool, showStatus string) error {
	file, err := os.Open(fileName)
	if err != nil {
		return fmt.Errorf("błąd podczas otwierania pliku %s: %v", fileName, err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		status := getStatus(line)
		if showStatus != "" && status != showStatus {
			continue
		}

		if noColor {
			fmt.Println(line)
		} else {
			color := getColor(status)
			fmt.Printf("%s%s%s\n", color, line, Reset)
		}
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("błąd odczytu pliku %s: %v", fileName, err)
	}

	return nil
}

func processInput(input io.Reader, noColor bool, showStatus string) error {
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := scanner.Text()
		status := getStatus(line)

		if showStatus != "" && status != showStatus {
			continue
		}

		if noColor {
			fmt.Println(line)
		} else {
			color := getColor(status)
			fmt.Printf("%s%s%s\n", color, line, Reset)
		}
	}
	if err := scanner.Err(); err != nil {
		return fmt.Errorf("błąd odczytu danych wejściowych: %v", err)
	}
	return nil
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
