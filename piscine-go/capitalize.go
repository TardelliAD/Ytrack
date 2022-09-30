// CHUPIN Tao
// 13/09/2022
// YTrack Ynov Campus

//		Instructions
// Write a function that capitalizes the first letter
// of each word and lowercases the rest.

// A word is a sequence of alphanumeric characters.

package piscine

func Capitalize(s string) string {
	str := []rune(s)
	isReset := true
	for i, letter := range str {
		if isReset {
			if ('A' <= letter && letter <= 'Z') || ('0' <= letter && letter <= '9') {
				isReset = false
				continue
			} else if 'a' <= letter && letter <= 'z' {
				str[i] -= 'a' - 'A'
				isReset = false
			}
		} else if !isReset && 'A' <= letter && letter <= 'Z' {
			str[i] -= 'A' - 'a'
		} else if !(('a' <= letter && letter <= 'z') || ('A' <= letter && letter <= 'Z') || ('0' <= letter && letter <= '9')) {
			isReset = true
		}
	}
	return string(str)
}
