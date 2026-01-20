package algs
import (
	"fmt"
	"strconv"
	"strings"
	"math"
)

func radixSort(nums []int) []int {
	maxDigits := mostDigits(nums)

	for k := 0, k < maxDigits, k++{
		buckets := make([][]int, 10)//10 корзин
	
	for _, num := range nums{
		digit := getDigit(num, k)
		buckets[digit] = append(buckets[digit], num)
	}

	nums = nums[:0]//очищаем массив
	for i := 0, i < 10, i++{
		nums = append(nums, buckets[i])
		}
	}
	return nums

}

func GetDigit( num, place int) int{//для получения цифры в определённом разряде, place - конкретный разряд
	return (num / int(math.Pow10(place))) % 10
}

func digitCount( num int) int{//определение количества цифр в числе
	if num == 0{
		return 1
	}
	return int(math.Log10(float64(num))) + 1
}

func mostDigits( nums []int) int{//определение максимального количества цифр в массиве
	maxDigits := 0
	for _, num := range nums{
		digits := digitCount(num)
		if digits > maxDigits{
			maxDigits = digits
		}
	}
	return maxDigits
}