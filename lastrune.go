package piscine

func LastRune(s string) rune {
	if len(s) > 0 {
		return rune(s[4])
	}
	return rune(0)
}
