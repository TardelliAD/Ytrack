package piscine

func FindNextPrime(nb int) int {
	var result int
	var i int
	if nb <= 1 {
		return 0
	} else {
		for i = 2; i <= nb/2; i++ {
			result = nb % i
			if result == 0 {
				nb += 1
			}
		}
		return nb
	}
}
