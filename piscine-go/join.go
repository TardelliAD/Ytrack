package piscine

func Join(strs []string, sep string) string {
	var newString string

	for i := 0; i < len(strs); i++ {
		newString += strs[i]

		if i != len(strs)-1 {
			newString += sep
		}
	}
	return newString
}
