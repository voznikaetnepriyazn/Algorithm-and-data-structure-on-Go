package algs

//n дипломов hхw, которые надо разместить на квадратной поверхности
//найти минимальную сторону квадрата
//мин значение(левая граница) - самая длинная сторона диплома
//макс значение(правая граница) - длина всех дипломов в один ряд
//на каждой итерации считать - кол-во дипломов в высоту и ширину - строки и столбцы
//middle / h(w) - кол-во возможных строк(столбцов)

func Search(h, w, n int) int {
	left := Max(w, h)
	right := left * n
	for left+1 < right {
		middle := (left + right) / 2
		res := (middle / w) * (middle / h) //кол-во дипломов, которое поместиться в сторону квадрата длиной middle
		if res < n {
			left = middle
		}
		right = middle
	}
	return right
}

func Max(w, h int) int {
	if w > h {
		return w
	}
	return h
}
