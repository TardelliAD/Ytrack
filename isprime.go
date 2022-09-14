package piscine

func IsPrime(nb int) bool {
	if nb == 1 || nb < 0 {
		return false
	}
	for tempo := nb - 1; tempo > 1; tempo-- {
		if nb%tempo == 0 {
			return false
		}
	}
	return true
}
