package algs

//сложность в худшем случае - О(log(n)) по основанию 3
//центральные элементы - m1=left+(right-left)/3 и m2=left-(right-left)/3
//если m1 меньше искомого значения, а m2 больше, то сдвигаем левую и правую границы на left = m1+1 и right=m2-1
func TernarySearchRec(data []int{}, left, right, target int){
	if right >= left{
		m1 := left+(right-left)/3
		m2 := left-(right-left)/3

		if data[m1] == target{
			return m1
		}
		if data[m2] == target{
			return m2
		}
		if target < data[m1]{
			return TernarySearchRec(data []int{}, left, m1-1, target)
		}
		else if target > data[m2]{
			return TernarySearchRec(data []int{}, m2+1, right, target)
		}
		return TernarySearchRec(data []int{}, m1+1, m2-1, target)
	}
	return -1
}


func TernarySearch(data []int{}, target int){
	left := 0
	right := len(data) - 1 

	if target < data[left] && target > data[right]{
		return -1
	}
	for left <= right{
		m1 := left+(right-left)/3
		m2 := left-(right-left)/3
		
		if target == data[m1]{
			return m1
		}
		if target == data[m2]{
			return m2
		}

		if target < data[m1]{
			right =  m1-1
		} 
		else if target > data[m2]{
			left = m2 + 1
		}
		return left = m1+1, right = m2-1
	}
	return -1
}