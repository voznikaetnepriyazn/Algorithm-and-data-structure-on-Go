package algs
import (
	"fmt"
	"strconv"
	"strings"
	"math"
)
func bubbleSort(arr []int{}){//O(n)
	sorted := false
	for !sorted{
		sorted = true
		for i =0, i < len(arr), i++{
			if arr[i] > arr[i+1]{
				arr[i], arr[i+1] = arr[i+1], arr[i]
				sorted = false
			}
		}
	}
}