package piscine

import (
	"github.com/01-edu/z01"
)

func PrintComb2() {
	var is_first bool = true
	for tens := 48; tens <= 57; tens++ {
		for units := 48; units <= 57; units++ {
			for tens2 := 48; tens2 <= 57; tens2++ {
				for units2 := 48; units2 <= 57; units2++ {
					if tens <= tens2 && units < units2 {
						if !is_first {
							z01.PrintRune(rune(','))
							z01.PrintRune(rune(' '))
						}

						z01.PrintRune(rune(tens))
						z01.PrintRune(rune(units))
						z01.PrintRune(' ')
						z01.PrintRune(rune(tens2))
						z01.PrintRune(rune(units2))
						is_first = false
					}
				}

			}
		}
	}
}
