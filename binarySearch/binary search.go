package algs

//только на отсортированном массиве
//нет смысла искать число, если элемент под нулевым индексом больше искомого значения или последний элемент массива меньше искомого

func binarySearchh(data []int, target int) int {

	//итеративный подход, меньше затрат памяти

	left := 0
	right := len(data) - 1

	if target < data[left] && target > data[right] {
		return -1
	}

	for left <= right {
		middle := (left + right) / 2
		if target == data[middle] {
			return middle
		}
		if target < data[middle] {
			right = middle - 1
		}
		left = middle + 1
	}
	return -1
}

//формула для избежания переполнения значений
//middle = left + (right - left) / 2

//сложность - на i-том проходе получаем: n / 2^i
//i = n / 2^i - i = log(n)
//O(log(n))

//необходима сортировка, поэтому n*log(n) - сложность сортировки, k - кол-во операций поиска
//n*k >= n*log(n) + k*log(n)

//левый бинарный поиск - первое вхождение

func leftBinarySearch(data []int, target int) int {
	left := 0
	right := target - 1

	for left+1 < right { //пока не останется 2 элемента
		middle := (left + right) / 2
		if data[middle] < target {
			left = middle
		}
		right = middle
	}

	if data[left] == target {
		return left
	}
	if data[right] == target {
		return right
	}
	return -1
}

//правый бинарный поиск - последнее вхождение

func rightBinarySearch(data []int, target int) int {
	left := 0
	right := target - 1

	for left+1 < right {
		middle := (left + right) / 2
		if data[middle] <= target {
			left = middle
		}
		right = middle
	}

	if data[right] == target {
		return right
	}
	if data[left] == target {
		return left
	}
	return -1
}
