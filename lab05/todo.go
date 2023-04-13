package main

import (
	"fmt"
	"os"
)

func main() {

}

func load_file(name string) {
	file, err := os.Open(name)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

}
