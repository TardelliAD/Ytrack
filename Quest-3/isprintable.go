package main

import "fmt"

func IsPrintable(s string) bool {
	for _, letter := range s {
		if !(letter >= 32 && letter <= 126) {
			return false
		}
	}
	return true
}
func main() {
	fmt.Println(IsPrintable("Hello"))
	fmt.Println(IsPrintable("Hello\n"))

}
