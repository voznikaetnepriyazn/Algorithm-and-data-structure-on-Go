package algs

//сколько последовательностей длины n, состоящих из 0 и 1, в котрых нет двух стоящих подряд единиц

func nnn(n int) any {
	dp := []int{1, 2}
	for i := 2; i < n; i++ {
		dp = append(dp, dp[i-1]+dp[i-2])
	}
	return dp[n]
}

//нет k стоящих подряд единиц
func kkk(n, k int) int {
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, k)
	}

	dp[0][0] = 1 //пустая последовательность

	for i := 1; i <= n; i++ {
		for j := 0; j < k; j++ {
			dp[i][0] += dp[i-1][j] //добавляем 0 - обнуляем счетчик единиц
		}
		for j := 0; j < k; j++ {
			dp[i][j] = dp[i-1][j-1] //добавляем 1 - если до этого было j-1 единиц подряд, теперь j
		}
	}

	total := 0
	for j := 0; j < k; j++ {
		total += dp[n][j]
	}
	return total
}
