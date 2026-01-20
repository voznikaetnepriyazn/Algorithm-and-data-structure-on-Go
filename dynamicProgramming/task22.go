package algs

//допрыгать до n за 1 или 2(3) прыжка - сколько способов существует
func Jump2(n int) int {
	dp2 := make([]int, n+1)
	dp2[0] = 0
	dp2[1] = 1
	for i := 2; i <= n; i++ {
		dp2[i] = dp2[i-1] + dp2[i-2]
	}
	return dp2[n]
}

func Jump3(n int) {
	dp := make([]int, n+1)
	dp[0] := 1
	dp[1] := 1
	dp[2] := 2
	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2] + dp[i-3]
	}
	return dp[n]
}

//какие значения хранить в dp
//определиться с рекуррентным соотношением
//определиться с базовыми кейсами где не работает рекуррентная формула
//определиться с порядком вычисления значений
//сформулировать ребования к ответу

//пусть кузнечик может прыгать на k шагов - на последний столбик может прыгнуть с n-k столбика
func Jump(n int, k int) int {
	dp := make([]int, n+1)
	dp[0] := 1

	for i := 1; i <= n; i++ {
		r := min(k, i)

		dp[i] := 0
		for j := 1; j <= r; j++ {
			dp[i] = dp[i] + dp[i-j] //вложенный цикл заменяет условие dp[n] = dp[n-1] + dp[n-2] + ... + dp[n-k]
		}
	}
}

//добавляем монетки - кузнечик может собирать и терять их. надо собрать их максимальное количество

func MaxCoins(n int, k int, coins []int) (int, int, []int) {
	dp := make([]int, n+1)

	//dp[i] - максимум монет до клетки i
	for i := 1; i <= n; i++ {
		for j := 1; j <= min(k, i); j++ {
			if dp[i] < dp[i-j]+coins[i-1] {
				dp[i] = dp[i-j] + coins[i-1]
			}
		}
	}
	maxCoinsCelected := dp[n]

	//восстановление пути прыжков
	jumps := []int{}
	i := n
	for i > 0 {
		for j := min(k, i); j > 0; j-- {
			if dp[i] == dp[i-j]+coins[i-1] {
				jumps = append(jumps, i)
				i -= j
				break
			}
		}
	}
	reverse(jumps)

	return maxCoinsCelected, len(jumps), jumps
}

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
