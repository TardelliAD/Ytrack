package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arguments := os.Args
	for _, value := range arguments[1:] {
		for _, value2 := range value {
			z01.PrintRune(value2)
		}
		z01.PrintRune('\n')
	}
}
