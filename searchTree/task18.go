package algs

func Maxx(str string) int {
	count := make(map[string]int)
	for i := 0; i < len(str); i++ {
		count[str[i]]++
	}
	max := 0
	for i, j := range count {
		if j > max {
			max = j
		}
	}
	return max
}
