package main

import "fmt"

func findnextprime(nb int) int {
	var result int
	var x int
	if nb <= 1 {
		return 0
	} else {
		for x = 2; x <= nb/2; x++ {
			result = nb % x
			if result == 0 {
				nb += 1
			}
		}
	}
	return nb

}
func main() {
	fmt.Println(findnextprime(5))
	fmt.Println(findnextprime(4))
	fmt.Println(findnextprime(2))
}
