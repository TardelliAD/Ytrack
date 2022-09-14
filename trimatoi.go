package piscine

func TrimAtoi(sentence string) int {
	sign := 1
	var result int
	var i int
	premier_chiffre := true
	for i < len(sentence) {
		if sentence[i] == '-' && premier_chiffre {
			sign = -1
		} else if sentence[i] > 48 && sentence[i] <= 57 {
			result = result*10 + int(sentence[i]-48)
			premier_chiffre = false
		}
		i++
	}
	return result * sign
}
