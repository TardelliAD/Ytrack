package piscine

func Makerange(min, max int) []int {
	var result []int
	if max < min {
		return result
	}
	result = make([]int, max-min)
	for i := 0; i < len(result); i++ {
		result[i] = min + i
	}
	return result
}
