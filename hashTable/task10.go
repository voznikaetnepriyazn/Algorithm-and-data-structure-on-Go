package algs

//найти 2 элемента в массиве, сумма которых равна target
func TwoSum(data []int, target int) (int, int) {
	cashe := make(map[int]int)
	for i := 0; i < len(data); i++ {
		//вычисляем возможный 2 элемент
		cashe[data[i]] = i
	}
	for i := 0; i < len(data); i++ {
		//вычисляем возможный 2 элемент
		diff := target - data[i]
		if _, exists := cashe[diff]; exists {
			return i, cashe[diff]
		}
	}
	return 0, 0
}
