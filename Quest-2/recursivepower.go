package main

import "fmt"

func recursivepower(nb int, power int) int {
	if power == 1 {
		return nb
	}
	if power > 1 {
		return nb * recursivepower(nb, power-1)
	}
	return 0
}

func main() {
	fmt.Println(recursivepower(4, 3))
}
