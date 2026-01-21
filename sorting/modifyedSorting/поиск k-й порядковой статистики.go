package algs

//найти k-тое по величине число в неотсортированном массиве O(n)

func Poisk(arr []int) int {
	pivot := median(arr)

	i, j := 0, len(arr)-1

	for i <= j {
		for arr[i] < pivot {
			i++
		}
		for arr[j] > pivot {
			j--
		}
		if i <= j {
			arr[i], arr[j] = arr[j], arr[i]
			i++
			j--
		}
	}

	var k int
	if len(arr[i:]) == k-1 {
		return i
	}

	if len(arr[i:]) >= k {
		Poisk(arr[i:])
	}
	Poisk(arr[:j+1])

	return k
}
