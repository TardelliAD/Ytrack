package main

import (
	"fmt"
)

func main() {
	fmt.Println(firstRune(""))
}

func firstRune(s string) rune {
	if len(s) > 0 {
		return rune(s[0])
	}
	return rune(0)
}
