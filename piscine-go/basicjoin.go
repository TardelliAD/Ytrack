package piscine

func BasicJoin(elems []string) string {
	var newString string
	for i := 0; i < len(elems); i++ {
		newString += elems[i]
	}
	return newString
}
