package piscine

func StrRev(s string) string {
	var reverseString string
	for letter := range s {
		reverseString = string(s[letter]) + reverseString
	}
	return reverseString
}
