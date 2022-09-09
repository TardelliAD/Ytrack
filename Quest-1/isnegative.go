package piscine

import "github.com/01-edu/z01"

func isnegative(nb int) {
	if nb < 0 {
		z01.PrintRune('T')

	} else {
		z01.PrintRune('F')
	}

	isnegative(7)

	z01.PrintRune('\n')
}
