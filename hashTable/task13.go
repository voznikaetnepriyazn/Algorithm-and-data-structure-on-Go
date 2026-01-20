package algs

func Find(n int, arr []int) int {
	sort1 := make(map[int]int) //[число]частота повторения
	k := n / 2
	for _, num := range arr {
		sort1[num]++
	}
	for num, i := range sort1 {
		if i > k {
			return num
		}
	}
	return -1
}
