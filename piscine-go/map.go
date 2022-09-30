// CHUPIN Tao
// 14/09/2022
// YTrack Ynov Campus

//		Instructions
// Write a function Map that, for an int slice, applies a function of this type func(int) bool
// on each element of that slice and returns a slice of all the return values.

package piscine

func Map(f func(int) bool, a []int) []bool {
	counter := 0
	for range a {
		counter++
	}
	result_a := make([]bool, counter)
	for i, val := range a {
		result_a[i] = f(val)
	}
	return result_a
}
