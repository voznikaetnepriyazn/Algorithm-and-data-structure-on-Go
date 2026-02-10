package algs

// mergeSort сортирует массив с использованием алгоритма слияния (O(n log n))
func mergeSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	mid := len(arr) / 2

	left := mergeSort(arr[:mid])

	right := mergeSort(arr[mid:])

	return merge(left, right)
}

// merge объединяет два отсортированных массива в один
func merge(a, b []int) []int {
	result := make([]int, 0, len(a)+len(b))
	i, j := 0, 0

	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			result = append(result, a[i])
			i++
		} else {
			result = append(result, b[j])
			j++
		}
	}

	result = append(result, a[i:]...)
	result = append(result, b[j:]...)

	return result
}
