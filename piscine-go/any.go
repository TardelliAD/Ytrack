package piscine

func Any(f func(string) bool, a []string) bool {
	isNumeric := false
	for _, val := range a {
		if f(val) {
			isNumeric = true
		}
	}
	return isNumeric
}
