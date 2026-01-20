package algs

//linear search - O(n)
//данные не отсортированны или поиск максимума или минимума
func linSearch(data []int, target int) int {
	for i := 0; i < len(data); i++ {
		if data[i] == target {
			return i
		}
	}
	return -1
}
