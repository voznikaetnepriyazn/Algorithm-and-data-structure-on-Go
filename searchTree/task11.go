package algs

//является ли дерево симметричным - лево и право зеркальны

//итеративный метод - обход в ширину

type TreeNode2 struct {
	data  int
	left  *TreeNode2
	right *TreeNode2
}

func IsSymmetric(root *TreeNode2) bool {
	if root == nil {
		return true
	}
	queue := []*TreeNode2{root}

	for len(queue) > 0 {
		queueLen := len(queue)
		for i := 0; i < queueLen; i++ {
			//логические операции над симметричными узлами
			if queue[i] == nil && queue[queueLen-i-1] == nil {
				continue
			}
			if queue[i] == nil || queue[queueLen-i-1] == nil {
				return false
			}
			if queue[i].data != queue[queueLen-i-1].data {
				return false
			}
			if queue[i] != nil {
				queue = append(queue, queue[i].left)
				queue = append(queue, queue[i].right)
			}
		}
		queue = queue[queueLen:] //удаление обработанных узлов
	}
	return true
}

//обход в глубину - рекурсивно

func dept(root *TreeNode2, res []int) []int {
	if root == nil {
		return res
	}

	dept(root.left, res)
	res = append(res, root.data)
	dept(root.right, res)

	return res
}

func IsSymmetricDFS(root *TreeNode2) bool {
	if root == nil {
		return true
	}
	res := make([]int, 0)
	data := []int{}
	data = dept(root, res)
	j := len(data) - 1
	for i := 0; i < len(data)/2; i++ {
		if data[i] != data[j] {
			return false
		}
		j--
	}
	return true
}
