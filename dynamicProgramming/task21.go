package algs

//Дана последовательность целых чисел и число K .
// Ваша задача — найти максимальное произведение любой подпоследовательности из K элементов этой последовательности.

func MaxProduct(N int, arr []int, k int) int {
	if k > n || k <= 0 {
		return 0
	}

	currentproduct := 1 //сюда вписываем результат умножения
	for _, i := range k {
		currentproduct = currentproduct * arr[i]
	}
	maxproduct := currentproduct

	for i := k; i < n; i++ { //перемещаем по всей последовательности
		if arr[i-k] == 0 { //пересчитываем если элемент равен нулю
			currentproduct = 1
			for j := i - k + 1; j < min(i+k, n); j-- { //не выходим за пределы k
				currentproduct = currentproduct * arr[j]
			}
		}
		currentproduct = (currentproduct / arr[i-k]) * arr[i] //удаляем из произведения текущий элемент и добавляем новый
		maxproduct = max(maxproduct, currentproduct)
	}

	return maxproduct
}
