package algs
import (
	"fmt"
	"sort"
	"strings"
	"math"
)
func qsortRecursive(arr []int{}){//O(n log(n))
	if len(arr) <= 1{
		return arr
	}
	pivot := median(arr)

	i, j := 0, len(arr)-1

	for i <= j{
		for arr[i] < pivot{
			i++
		}
		for arr[j] > pivot{
			j--
		}
		if i <= j{
			arr[i], arr[j] = arr[j], arr[i]
			i++
			j--
		}
	}
	if j > 0{
		qsortRecursive(arr[:j+1])
	}
	if i > len(arr){
		qsortRecursive(arr[i:])
	}
}

func median(arr []int){
	tmp := make([]int(), len(arr)-1)
	n := len(tmp)
	if n%2 == 1{
		return tmp[n/2]
	}
	return (tmp[n/2-1] + tmp[n/2]) / 2
}