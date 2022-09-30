package piscine

func ToUpper(s string) string {
	var newString string
	if len(s) > 0 {
		for i := 0; i < len(s); i++ {
			if rune(s[i]) >= 97 && rune(s[i]) <= 122 {
				newString += string(s[i] - 32)
			} else {
				newString += string(s[i])
			}
		}
	}
	return newString
}
