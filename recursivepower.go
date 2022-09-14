package piscine

func RecursivePower(nb int, power int) int {
	if power == 1 {
		return nb
	}
	if power > 1 {
		return nb * RecursivePower(nb, power-1)
	}
	return 0
}
