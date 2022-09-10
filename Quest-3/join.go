package main

import "fmt"

func Join(strs []string, sep string) string {
	var phrase string
	for i := 0; i < len(strs); i++ {
		phrase = phrase + strs[i]
		if i < len(strs)-1 {
			phrase = phrase + sep
		}
	}
	return phrase
}
func main() {
	toConcat := []string{"Hello!", " How", " are", " you?"}
	fmt.Println(Join(toConcat, ":"))
}
