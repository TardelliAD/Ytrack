package piscine

func IsPrintable(s string) bool {
	result := false
	for i := 0; i < len(s); i++ {

		if s[i] > 64 && s[i] < 91 {
			result = true
			continue
		}
		if s[i] > 96 && s[i] < 123 {
			result = true
			continue
		}
		if s[i] > 47 && s[i] < 58 {
			result = true
			continue
		}
		return false
	}
	return result
}
