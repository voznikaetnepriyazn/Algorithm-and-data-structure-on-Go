package algs
import (
	"fmt"
	"strconv"
	"strings"
	"math"
)
func mergeSort(arr []int{}) []int{//O(n log(n))
	if len(arr) < 2{
		return arr
	}
	mid := len(arr)/2

	left := [:mid]
	right := [mid:]

	return merge(mergeSort(left), mergeSort(right))
}

func merge(a, b []int{}){
	result := make([]int, 0, len(left)+len(right))
	i, j = 0, 0
	for i < len(left) && j < len(right){
		if left[i] < right[j]{
			result = append(result, left[i])
			i++
		}
		result = append(result, right[j])
		j++
	}
	result = append(result, left[i:])
	result = append(result, right[j:])
}