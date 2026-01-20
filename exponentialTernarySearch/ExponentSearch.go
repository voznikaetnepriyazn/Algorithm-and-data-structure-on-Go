package algs
import (
	"fmt"
	"strconv"
	"strings"
	"math"
)
//для уточнения диапазона поиска
//O(log(m))- m индекс элемента, который нужно найти
func ExponentSearch(data []int{}, target int){
	border := 1
	if data[border] < target{
		border = border*2
	}
	if border > len(data){
		left := border/2
		right := len - 1
		middle := (left + right)/2
		if data[middle] == target{
			return middle
		}
		if data[middle] < target{
			left = middle + 1
		}
		right = middle - 1
	}
	if border > target{
		left := border/2
		right := border
		middle := (left + right)/2
		if data[middle] == target{
			return middle
		}
		if data[middle] < target{
			left = middle + 1
		}
		right = middle - 1
	}
}

//найти отрезок массива, в котором располагается число target
func expSearch( data []int{}, target int){
	border := 1
	lastElement:= len(data) - 1

	for border < lastElement && data[border] <= target{
		if data[border] == target{
			return [border, border*2]
		}
		border = border*2
		if border > lastElement{
			return [border/2, lastElement]
		}
	}

	return [border / 2, border]
}