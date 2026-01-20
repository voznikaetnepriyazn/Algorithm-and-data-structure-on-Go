package algs

func Delete(n int, arr []int, element int) []int {
	/*ive := o
	for i = 0; i < n; i++{
		if arr[i] != element{
			arr[ive] = arr[i]
			ive ++
		}
	}
	var res string
	i
	for _, num := range arr{
		if arr[i] != element{
			res = append(res, strconv.Itoa(num))
		}
	}*/

	arr1 := make([]int, 0)
	for _, num := range arr {
		if num != element {
			arr1 = append(arr1, num)
		}
	}
	return arr1

}
