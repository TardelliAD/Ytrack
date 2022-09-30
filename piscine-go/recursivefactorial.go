package piscine

func RecursiveFactorial(index int) int {
	if index == 1 {
		return 1
	}
	if index > 1 {
		return index * RecursiveFactorial(index-1)
	}
	return 0
}
