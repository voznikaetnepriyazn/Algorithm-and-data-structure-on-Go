package algs

import (
	"strconv"
)

func SortGrades(n int, arr []int) []string {
	a := 0
	for i := 0; i < n; i++ {
		if arr[i] != 0 {
			arr[i] = arr[a]
			a++
		}
	}
	for i := a; i < n; i++ {
		arr[i] = 0
	}
	var res []string
	for _, num := range arr {
		res = append(res, strconv.Itoa(num))
	}
	return res
}
