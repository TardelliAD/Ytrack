// CHUPIN Tao
// 13/09/2022
// YTrack Ynov Campus

//		Instructions
// Write a function that separates the words
// of a string and puts them in a string slice.

// The separators are spaces, tabs and newlines.

package piscine

func SplitWhiteSpaces(s string) []string {
	var slice []string
	var word string
	for i := 0; i < len(s); i++ {
		if s[i] == 32 || s[i] == 9 || s[i] == 10 {
			slice = append(slice, word)
			word = ""
			continue
		}
		word += string(s[i])
	}
	slice = append(slice, word)
	return slice
}
