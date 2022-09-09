package main

import (
	"fmt"
)

func main() {
	fmt.Println(lastRune("Hello"))
}

func lastRune(s string) rune {
	if len(s) > 0 {
		return rune(s[4])
	}
	return rune(0)
}
