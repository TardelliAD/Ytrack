package piscine

import "github.com/01-edu/z01"

func PrintComb() {
	var pute bool = true
	for centaine := 48; centaine < 56; centaine++ {
		for dizaine := 49; dizaine < 57; dizaine++ {
			for unite := 50; unite < 58; unite++ {
				if centaine < dizaine && dizaine < unite {
					if !pute {
						z01.PrintRune(',')
						z01.PrintRune(' ')
					}
					z01.PrintRune(rune(centaine))
					z01.PrintRune(rune(dizaine))
					z01.PrintRune(rune(unite))
					pute = false

				}

			}
		}
	}
	z01.PrintRune('\n')
}
