package main

import "fmt"

func IsUpper(s string) bool {
	for _, letter := range s {
		if !(letter >= 'A' && letter <= 'Z') {
			return false
		}
	}
	return true
}
func main() {
	fmt.Println(IsUpper("HELLO"))
	fmt.Println(IsUpper("HELLO!"))

}
