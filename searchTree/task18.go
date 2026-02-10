package algs

func Maxx(str string) int {
	count := make(map[rune]int)
	for _, char := range str {
		count[char]++
	}

	max := 0

	for _, j := range count {
		if j > max {
			max = j
		}
	}
	return max
}
