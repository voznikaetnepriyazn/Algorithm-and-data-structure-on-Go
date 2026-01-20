package algs
import (
	"fmt"
	"strconv"
	"strings"
	"math"
)

//собираем монетки в 2мерном пространстве
//двигаемся только вправо или вниз
//движение по первой горизонтали - dp[0][k]=dp[0][k-1]+Coins[0][k]
//по вертикали - dp[k][0]=dp[k-1][0]+Coins[k][0]

func Tortill(coins [][]int) [][]int{
	if len(coins) == 0 || len(coins[0]) == 0{
		return [][]int{}
	}
	n := len(coins)
	m := len(coins[0])

	dp := make([][]int, n)
	for _, i := range dp{
		dp[i] = make([]int, m)
	}

	prev := make([][]int, n)//храним маршрут
	for _, i := range prev{
		prev[i] = make([]int, m)
	}

	for i := 0; i < n; i++{
		for j := 0; j < m; j++{
			if i == 0 && j == 0{
				dp[0][0] = coins[0][0]
				prev[0][0] = -1//предыдущей клетки нет
			}
			else if i == 0{
				dp[0][j] = dp[0][j-1] + coins[0][j]
				prev[0][j] = 0//слева пришли
			}
			else if j == 0{
				dp[i][0] = dp[i-1][0] + coins[i][0]
				prev[i][0] = 1//сверху пришли
			}
			dp[i][j] = max(dp[i-1][j], dp[i][j-1]) + coins[i][j]
			if dp[i-1][j] > dp[i][j-1]{
				prev[i][j] = 1
			}
			prev[i][j] = 0
		}
	}
	return dp, prev[i]
	
}