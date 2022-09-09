// Write a function that returns the nth rune of a string. If not possible, it returns 0.

package main

import (
	"fmt"
)

func main() {

	fmt.Println(nrune("Hello", 4))
	fmt.Println(nrune("Hello", 1))
}

func nrune(s string, n int) rune {
	if len(s) > 0 {
		return rune(s[n-1])
	}
	return rune(0)

}
