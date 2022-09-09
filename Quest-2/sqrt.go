package main

import "fmt"

func Sqrt(nb int) int {
	var result int
	if nb == 0 {
		return 0
	}
	if nb > 1 {
		for mult := 1; result < nb; mult++ {
			result = mult * mult
			if result == nb {
				return mult
			}
		}
	}
	return 0
}
func main() {
	fmt.Println(Sqrt(10))
}
