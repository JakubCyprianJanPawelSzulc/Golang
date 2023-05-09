package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

var RANGE int = 1000
var COLUMN int = 0
var FORMAT string = "%12.4f%12.4f\n"

func main() {
	files := AnalyzeOPT(os.Args)
	var lines []string = make([]string, 0)

	fmt.Println(files)
	if len(files) == 0 {
		files = append(files, "-")
	}

	for _, file := range files {
		var f *os.File
		if file == "-" {
			f = os.Stdin
		} else {
			var err error
			f, err = os.Open(file)
			defer f.Close()
			if err != nil {
				fmt.Println("[w] błąd otwarcia pliku:", file, err)
				continue
			}

			scanner := bufio.NewScanner(f)
			for scanner.Scan() {
				lines = append(lines, scanner.Text())
			}

			if len(lines) > 0 {
				var hist map[int]int = make(map[int]int)

				for _, line := range lines {
					parts := strings.Fields(line)
					var part string
					var val int
					if len(parts) > COLUMN {
						part = parts[COLUMN]
						if v, err := strconv.ParseFloat(part, 64); err == nil {
							val = int(v)
						}
					}
					val = val / RANGE
					hist[val]++
				}

				keys := make([]int, 0, len(hist))
				for k := range hist {
					keys = append(keys, k)
				}
				sort.Ints(keys)
				for _, k := range keys {
					v := hist[k]
					fmt.Printf(FORMAT+"\n", float64(k*RANGE), float64(v))
				}
			} else {
				fmt.Println("[w] Brak wierszy do liczenia histogramu")
			}
		}
		fmt.Println("plik: ", file)
	}
}

func AnalyzeOPT(args []string) []string {

	var pliki []string = make([]string, 0)

	for i := 1; i < len(args); i++ {
		// fmt.Println("opcja:",i,args[i])
		opt := args[i]

		if len(opt) < 2 {
			pliki = append(pliki, opt)
			continue
		}

		switch opt[:2] {
		case "-f":
			if len(opt) > 2 {
				FORMAT = opt[2:]
				continue
			} else {
				if i+1 < len(args) {
					FORMAT = args[i+1]
					i++
					continue
				} else {
					fmt.Println("[w] opcja -f MUSI mieć argument")
				}
			}
		case "-c":
			if len(opt) > 2 {
				COLUMN = ustawopt(opt[2:], COLUMN, "błąd w opcji -c")
			} else {
				if i+1 < len(args) {
					COLUMN = ustawopt(args[i+1], COLUMN, "błąd w opcji -c")
					i++
				} else {
					fmt.Println("[w] opcja -c wymaga argumentu")
				}
			}
		case "-r":
			if len(opt) > 2 {
				RANGE = ustawopt(opt[2:], RANGE, "błąd w opcji -r")
			} else {
				if i+1 < len(args) {
					RANGE = ustawopt(args[i+1], RANGE, "błąd po opcji -r")
					i++
				} else {
					fmt.Println("[w] opcja -r wymaga argumentu")
				}
			}
		default:
			if opt[:1] == "-" {
				fmt.Println("[w] opcja bez sensu!", opt)
			} else {
				pliki = append(pliki, opt)
			}
		}
	}
	// fmt.Println("RANGE:",RANGE)
	// fmt.Println("COLUMN:",COLUMN)
	// fmt.Println("FORMAT:",FORMAT)
	return pliki
}

func ustawopt(opt string, def int, e string) int {
	if v, err := strconv.Atoi(opt); err == nil {
		return v
	} else {
		fmt.Println("[w]", e, err)
		return def
	}
}
