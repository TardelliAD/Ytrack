package piscine

func IsUpper(s string) bool {
	if len(s) > 0 {
		for i := 0; i < len(s); i++ {
			if rune(s[i]) >= 65 && rune(s[i]) <= 90 {
				continue
			} else {
				return false
			}
		}
		return true
	}
	return false
}
