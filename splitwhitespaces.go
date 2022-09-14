package piscine

func SplitWhiteSpaces(s string) []string {
	var arr []string
	var mot string
	for i := 0; i < len(s); i++ {
		if s[i] == 32 || s[i] == 9 || s[i] == 10 {
			arr = append(arr, mot)
			mot = ""
		} else {
			mot = mot + string(s[i])
		}

	}
	arr = append(arr, mot)
	return arr
}
