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
	for row := 0; row < n; row++ {
		currentRow := make([]int)     //массив для хранения строк
		for _, col := range row + 1 { //цикл по столбцам для кажой ячейки
			currentRow = append(currentRow, PascalTriangle(row, col))
		}
		dp = append(dp, currentRow)
	}
}

func PascalTriangleIteraive(row, col int) {
	for i := 1; i <= n; i++ {
		tmp := make([]int)
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
