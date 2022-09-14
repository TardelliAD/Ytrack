package piscine

import "fmt"

func ForEach(f func(int), a []int) {
	for _, nbr := range a {
		fmt.Print(nbr)
	}
}
