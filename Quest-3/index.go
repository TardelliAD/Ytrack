package main

import "fmt"

func Index(s string, toFind string) int {
	mot := []rune(s)
	toFindword := []rune(toFind)
	TailleMot := len(mot)
	TailleFind := len(toFindword)
	for i := 0; i <= TailleMot-TailleFind; i++ {
		if toFind == s[i:i+TailleFind] {
			return (i)
		}
	}
	return -1
}
func main() {
	fmt.Println(Index("Hello!", "l"))
	fmt.Println(Index("Salut!", "alu"))
	fmt.Println(Index("Ola!", "hOl"))
}
