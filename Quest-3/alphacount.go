// Write a function that counts the letters of a string and returns the count.

package main

import "fmt"

func AlphaCount(s string) int {
	alphacount := 0
	for i := 0; i < len(s); i++ {
		if s[i] >= 'a' && s[i] <= 'z' || s[i] >= 'A' && s[i] <= 'Z' {
			alphacount++
		}
	}
	return alphacount
}

func main() {
	s := "Hello, World! /"
	nb := AlphaCount(s)
	fmt.Println(nb)
}
