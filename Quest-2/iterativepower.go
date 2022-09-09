package main

import "fmt"

func iterativepower(nb int, power int) int {
	result := 1
	for i := power; i > 0; i-- {
		result *= nb

	}
	return result

}

func main() {
	fmt.Println(iterativepower(4, 3))
}
