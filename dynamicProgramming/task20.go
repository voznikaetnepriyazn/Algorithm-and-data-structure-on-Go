package algs

//рассчитать треугольник паскаля
var n int
var dp [][]int

func PascalTriangle(row, col int) int {
	//если это вершина треугольника или элемент его стороны, возвращаем 1
	if col == 0 || row == col {
		return 1
	}
	return PascalTriangle(row-1, col-1) + PascalTriangle(row-1, col)
}

func BuildPascalTriangle(n int) [][]int {
	dp := make([][]int, n)

	for row := 0; row < n; row++ {
		currentRow := make([]int, 0)      //массив для хранения строк
		for col := 0; col <= row; col++ { //цикл по столбцам для кажой ячейки
			currentRow = append(currentRow, PascalTriangle(row, col))
		}
		dp = append(dp, currentRow)
	}
	return dp
}

func PascalTriangleIteraive(row, col int) {
	for i := 1; i <= n; i++ {
		tmp := make([]int, 0)
		for j := 1; j <= i; j++ {
			tmp = append(tmp, 1)
		}
		dp = append(dp, tmp)
	}

	//заполняем массив значениями
	for row := 1; row < n; row++ {
		for col := 1; col < row; col++ {
			dp[row][col] = dp[row-1][col-1] + dp[row-1][col]
		}
	}
}
