// CHUPIN Tao
// 12/09/2022
// YTrack Ynov Campus

//		Instructions
// Write a function that returns true if the string passed
// as a parameter contains only printable characters, otherwise, returns false.

package piscine

import (
	"github.com/01-edu/z01"
)

func PrintWordsTables(a []string) {
	for _, word := range a {
		for _, letter := range word {
			z01.PrintRune(letter)
		}
		z01.PrintRune('\n')
	}
}
