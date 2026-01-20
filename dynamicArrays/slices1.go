package algs

//обьединение 2ух отсортированных по возрастанию массивов
func MergeSortedArrays(arr1 []int, arr2 []int) []int {
	margedArr := make([]int, len(arr1)+len(arr2))
	i := 0
	j := 0
	for i < len(arr1) && j < len(arr2) {
		if arr1[i] < arr2[j] {
			margedArr = append(margedArr, arr1[i])
			i++
		}
		margedArr = append(margedArr, arr2[j])
		j++
	}
	margedArr = append(margedArr, arr1[i:]...)
	margedArr = append(margedArr, arr2[j:]...)
	return margedArr
}
