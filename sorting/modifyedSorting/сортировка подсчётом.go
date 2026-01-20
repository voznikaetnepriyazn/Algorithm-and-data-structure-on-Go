package algs

/*func countSort(arr []int){
	s := Max(arr)
	arr1 := make([]int, s)
	for _, i range arr{
		arr1[i]++
	}
	arr2 := make([]int, len(arr))
	for _, j range arr1{

	}
}*/

func countingSort(arr []int) []int {

	if len(arr) == 0 {
		return arr
	}

	min, max := arr[0], arr[0]
	for _, v := range arr {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}

	count := make([]int, max-min+1)

	for _, v := range arr {
		count[v-min]++ //вычитаем мин, чтобы можно было работать с отрицательными числами
	}

	index := 0
	sorted := make([]int, len(arr))
	for i, c := range count {
		for c > 0 { //только при счетчике больше нуля
			sorted[index] = i + min
			index++
			c--
		}
	}
	return
}
