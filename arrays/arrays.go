package algs

//массивы
//О(n) - вставка и удаление элемента в середину массива

func TwoSum(arr []int, sum int) (int, int) { //O(n)
	left := 0
	right := len(arr) - 1
	for left != right {
		tmp := arr[left] + arr[right]
		if tmp == sum {
			return arr[left], arr[right]
		}
		if tmp < sum {
			left++
			continue
		}
		right--
	}

	return arr[left], arr[right]
}

func ReverseArray(arr []int) []int {
	left := 0
	right := len(arr) - 1
	for left < right {
		arr[left], arr[right] = arr[right], arr[left]
		left++
		right--
	}
	return arr
}
