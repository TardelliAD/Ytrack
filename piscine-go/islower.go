package piscine

func IsLower(s string) bool {
	if len(s) > 0 {
		for i := 0; i < len(s); i++ {
			if rune(s[i]) >= 97 && rune(s[i]) <= 122 {
				continue
			} else {
				return false
			}
		}
		return true
	}
	return false
}
