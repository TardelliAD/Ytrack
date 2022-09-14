package piscine

func NRune(s string, n int) rune {
	if len(s) > 0 {
		return rune(s[n-1])
	}
	return rune(0)

}
