package algs

func copyTime(n, x, y int) int {
	left := 0
	right := (n - 1) * Max(x, y) //1 лист необходимо напечатать за минимально возможное время
	for left+1 < right {
		middle := (left + right) / 2
		if middle/x+middle/y < n-1 { //сколько листов напечатает за время middle
			left = middle
		}
		right = middle
	}
	return right + Min(x, y) //1 копия делается на самом быстром ксероксе
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
