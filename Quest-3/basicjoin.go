package main

import "fmt"

func BasicJoin(elems []string) string {
	phrase := ""
	for _, taille := range elems {
		phrase = phrase + taille
	}
	return phrase
}
func main() {
	elems := []string{"Hello!", " How", " are", " you?"}
	fmt.Println(BasicJoin(elems))
}
