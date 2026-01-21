package algs

func selectionSort(arr []int) { //O(n^2)
	l := len(arr)
	for i := 0; i <= l-1; i++ {
		min := i
		for j := i + 1; j <= l; j++ {
			if arr[j] < arr[min] {
				min = j
			}
		}
		arr[i], arr[min] = arr[min], arr[i]
	}
}
