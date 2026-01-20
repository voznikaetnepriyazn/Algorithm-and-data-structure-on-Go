package algs
import (
	"fmt"
	"strconv"
	"strings"
	"math"
)
func insertionSort(arr []int{}){//O(n)
	for i = 1, i < len(arr), i++{
		j=i
		for j >0{
			if arr[j-1] > arr[j]{
				arr[j-1], arr[j] = arr[j], arr[j-1]
			}
			j--
		}
	}
}