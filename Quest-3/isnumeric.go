package main

import "fmt"

func IsNumeric(s string) bool {
	for _, letter := range s {
		if !(letter >= '0' && letter <= '9') {
			return false
		}
	}
	return true
}
func main() {
	fmt.Println(IsNumeric("010203"))
	fmt.Println(IsNumeric("01,02,03"))
}
