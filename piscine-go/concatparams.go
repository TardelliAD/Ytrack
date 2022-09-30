package piscine

func ConcatParams(args []string) string {
	var concatPara string
	for i := 0; i < len(args); i++ {
		concatPara += args[i]
		concatPara += "\n"
	}
	return concatPara
}
