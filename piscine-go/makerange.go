package piscine

func MakeRange(min, max int) []int {
	var tab []int
	nbAdd := min
	capTab := max - min

	if max < min {
		return tab
	}
	tab = make([]int, capTab)
	for i := 0; i < capTab; i++ {
		tab[i] = nbAdd
		nbAdd++
	}
	return tab
}
