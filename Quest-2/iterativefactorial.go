package main

import "fmt"

func x(index int) int {

	result := 1
	for i := 1; i < index+1; i++ {
		result = result * i

	}
	fmt.Println(result)
	return result
}

func main() {
	x(5)
}
