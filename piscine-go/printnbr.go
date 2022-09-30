package piscine

import "github.com/01-edu/z01"

func PrintNbr(nb int) {
	if nb < 0 {
		z01.PrintRune('-')
		nb = -nb
	}
	if nb > 9 {
		PrintNbr(nb / 10)
	}

	z01.PrintRune(rune(nb%10 + 48))

}
