package piscine

func IsPrime(nb int) bool {
	var result int
	var i int
	if nb <= 1 {
		return false
	} else {
		for i = 2; i <= nb/2; i++ {
			result = nb % i
			if result == 0 {
				return false
			}
		}
		return true
	}
}
