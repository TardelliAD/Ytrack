package piscine

func AppendRange(min, max int) []int {
	var tabEntier []int
	for i := min; i < max; i++ {
		tabEntier = append(tabEntier, min)
		min++
	}
	return tabEntier
}
