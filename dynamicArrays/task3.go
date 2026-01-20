package algs

func ADd(n int, arr []int) int {
	max := 0
	for _, num := range arr {
		if num > max {
			max = num
		}
	}
	return max
}
