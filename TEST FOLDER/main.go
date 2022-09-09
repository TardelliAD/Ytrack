package main

import "fmt"

func x(index int) int {

	if index == 1 {
		return 1
	}
	if index > 1 {
		return index * x(index-1)
	}
	return 0

}

func main() {

	result := x(5)
	fmt.Println(result)

}
