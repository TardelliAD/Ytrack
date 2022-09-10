// Write a function that transforms numbers within a string, into an int.

// If the - sign is encountered before any number it should determine the sign of the returned int.

// This function should only return an int. In the case of an invalid input, the function should return 0.

// Note: There will never be more than one sign in a string in the tests.

package main

import "fmt"

func TrimAtoi(s string) int {
	var sign int
	var result int
	var i int
	for i < len(s) {
		if s[i] == '-' {
			sign = -1
		} else if s[i] == '+' {
			sign = 1
		} else if s[i] >= '0' && s[i] <= '9' {
			result = result*10 + int(s[i]-'0')
		}
		i++
	}
	return result * sign
}
func main() {
	fmt.Println(TrimAtoi("12345"))
	fmt.Println(TrimAtoi("str123ing45"))
	fmt.Println(TrimAtoi("012 345"))
	fmt.Println(TrimAtoi("Hello World!"))
	fmt.Println(TrimAtoi("sd+x1fa2W3s4"))
	fmt.Println(TrimAtoi("sd-x1fa2W3s4"))
	fmt.Println(TrimAtoi("sdx1-fa2W3s4"))
	fmt.Println(TrimAtoi("sdx1+fa2W3s4"))
}
