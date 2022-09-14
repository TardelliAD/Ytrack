package piscine

import (
	"fmt"
)

func AppendRange(min, max int) []int {
	var result []int

	for i := min; i < max; i++ {
		result = append(result, i)
		if result == nil {
			fmt.Println("nil!")
		}
	}
	return result
}
