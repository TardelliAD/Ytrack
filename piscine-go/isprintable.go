// CHUPIN Tao
// 12/09/2022
// YTrack Ynov Campus

//		Instructions
//Write a function that returns true if the string passed as
//a parameter contains only printable characters, otherwise, returns false.

package piscine

func IsPrintable(s string) bool {
	if len(s) > 0 {
		for i := 0; i < len(s); i++ {
			if rune(s[i]) >= 48 && rune(s[i]) <= 57 {
				continue
			} else {
				return false
			}
		}
		return true
	}
	return false
}
