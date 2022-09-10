package main

import "fmt"

func IsLower(s string) bool {
	for _, letter := range s {
		if !(letter >= 'a' && letter <= 'z') {
			return false
		}
	}
	return true
}
func main() {
	fmt.Println(IsLower("hello"))
	fmt.Println(IsLower("hello!"))

}
