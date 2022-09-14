package piscine

func Split(s, sep string) []string {
	var new_string []string
	if len(s) == 0 || len(sep) == 0 || len(s) < len(sep) {
		return new_string
	}
	var temp string
	for i := 0; i < len(s); i++ {
		if s[i] == sep[0] {
			if Compare(s[i:i+len(sep)], sep) == 0 {
				new_string = append(new_string, temp)
				temp = ""
				i += len(sep) - 1
				continue
			}
		}
		temp += string(s[i])
	}
	new_string = append(new_string, temp)
	return new_string
}
