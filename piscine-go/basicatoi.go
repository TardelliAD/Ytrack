package piscine

func BasicAtoi(s string) int {
	sValue := []byte(s)
	count := 0
	for range sValue {
		count += 1
	}
	result := 0
	pow := 1
	for i := count - 1; i >= 0; i-- {
		result += ((int(sValue[i]) - 48) * pow)
		pow *= 10
	}
	return result
}
