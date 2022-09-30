package piscine

import (
	"github.com/01-edu/z01"
)

func Printcomb() {
	var is_first bool = true
	for hundreds := 48; hundreds <= 57; hundreds++ {
		for tens := 48; tens <= 57; tens++ {
			for units := 48; units <= 57; units++ {
				if hundreds < tens && tens < units {
					if !is_first {
						z01.PrintRune(rune(','))
						z01.PrintRune(rune(' '))
					}
					z01.PrintRune(rune(hundreds))
					z01.PrintRune(rune(tens))
					z01.PrintRune(rune(units))
					is_first = false
				}
			}
		}
	}
}
