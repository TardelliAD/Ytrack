package piscine

func IsNumeric(s string) bool {
	if len(s) > 0 {
		for i := 0; i < len(s); i++ {
			if rune(s[i]) >= 48 && rune(s[i]) <= 57 {
				continue
			} else {
				return false
			}
		}
		return true
	}
	return false
}
