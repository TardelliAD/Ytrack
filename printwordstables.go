package piscine

import (
	"github.com/01-edu/z01"
)

func PrintWordsTables(a []string) {
	var resultat string
	for i := range a {
		resultat += a[i]
		if i < len(a) {
			resultat += "\n"
		}
	}
	for i := range resultat {
		z01.PrintRune(rune(resultat[i]))
	}
}
func SplitWhiteSpaceses(s string) []string {
	var arr []string
	var mot string
	for i := 0; i < len(s); i++ {
		if s[i] == 32 || s[i] == 9 || s[i] == 10 {
			arr = append(arr, mot)
			mot = ""
		} else {
			mot = mot + string(s[i])
		}

	}
	arr = append(arr, mot)
	return arr
}
