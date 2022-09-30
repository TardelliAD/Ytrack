package piscine

func Index(s string, toFind string) int {
	for i := 0; i < len(s); i++ {
		if string(s[i]) == toFind {
			return i
		}
	}
	return 0
}
