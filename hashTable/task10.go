package algs
import (
	"fmt"
	"strconv"
	"strings"
	"math"
)
//найти 2 элемента в массиве, сумма которых равна target
func TwoSum(data []int, target int){
	cashe := make(map[int]int)//[array_element]index_of_array_element
	for i := 0, i < len(data), i++{
		//вычисляем возможный 2 элемент
		cashe[data[i]] = i
	}
	for i := 0, i < len(data), i++{
		//вычисляем возможный 2 элемент
		diff := target - data[i]
		if _, exists := cashe[diff]; exists{
			return [i, cashe[diff]]
		}
	}
}