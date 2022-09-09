/*
	Write a function that returns true if the int passed as parameter is a prime number. Otherwise it returns false.

The function must be optimized in order to avoid time-outs with the tester.

(We consider that only positive numbers can be prime numbers)

(We also consider that 1 is not a prime number)
*/
package main

import "fmt"

func isprime(n int) bool {
	if n == 1 || n < 0 {
		return false
	}
	for tempo := n - 1; tempo > 1; tempo-- {
		if n%tempo == 0 {
			return false
		}

	}
	return true
}
func main() {
	fmt.Println(isprime(5))
	fmt.Println(isprime(4))
	fmt.Println(isprime(2))
	fmt.Println(isprime(15))
	fmt.Println(isprime(13))
}
