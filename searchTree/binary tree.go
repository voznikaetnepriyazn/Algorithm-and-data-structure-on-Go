package algs

import (
	"fmt"
)

//дерево
//состоит из нод

// бинарное дерево - у каждой ноды не более 2ух потомков
// меньше проверок при рекурсивном обходе 0 < n < 2
type NodeOfTree struct {
	data  int
	left  *NodeOfTree
	right *NodeOfTree
}

//бинарное дерево поиска О(log n)
//поддерево каждого узла - тоже бинарное дерево поиска
//у каждой вершины в левом поддереве все элементы не меньше, чем в самой вершине
//в правом поддереве - больше

func Minimum(root *NodeOfTree) *NodeOfTree {
	if root.left == nil {
		return root
	}
	return Minimum(root.left)
}

func MinimumNonRec(root *NodeOfTree) *NodeOfTree {
	for root.left == nil {
		root = root.left
	}
	return root
}

func Maximum(root *NodeOfTree) *NodeOfTree {
	if root.right == nil {
		return root
	}
	return Maximum(root.right)
}
func MaximumNonRec(root *NodeOfTree) *NodeOfTree {
	for root.right == nil {
		root = root.right
	}
	return root
}

// поиск О(log n) в сбалансированном, в худшем случае - О(n)
func Search(root *NodeOfTree, target int) *NodeOfTree {
	if root == nil {
		return nil
	}
	if root.data == target {
		return root
	}
	if target < root.data {
		return Search(root.left, target)
	}
	return Search(root.right, target)
}

//сбалансированное дерево - высоты 2ух поддеревьев каждого из узлов отличаются не более чем на единицу

func Insert(root *NodeOfTree, target int) *NodeOfTree {
	var data int

	if root == nil {
		return &NodeOfTree{
			data: data,
		} //конструктор узла
	}
	if target < root.data {
		root.left = Insert(root.left, data)
	}
	root.right = Insert(root.right, data)

	return root
}

// удаление листа, удаление ноды с одним потомком, удаление ноды с 2мя узлами - О(log n)
func Delete(root *NodeOfTree, target int) *NodeOfTree {
	if root == nil {
		return root
	}

	//поиск ключа
	if target < root.data {
		root.left = Delete(root.left, target)
	}
	if target > root.data {
		root.right = Delete(root.right, target)
	}

	//узел является листом - без потомков
	if root.left == nil && root.right == nil {
		root = nil
	}
	return root

	//узел имеет одного потомка
	if root.left == nil {
		root = root.right
		return root
	}
	if root.right == nil {
		root = root.left
		return root
	}

	//2 потомка
	minVal := Minimum(root.right)
	root.data = minVal.data                      //вместо удаляемого значения записываем найденное минимальное
	root.right = Delete(root.right, minVal.data) //удаление минимального значения с его прежней позиции

	return root
}

// обход в ширину - сверху вниз, слева направо, итеративный алгоритм обхода
func breadthTree(root *NodeOfTree) []int {
	if root == nil {
		return []int{}
	}

	var result []int
	currentlevel := []*NodeOfTree{root} //срез, здесь корень дерева

	for len(currentlevel) > 0 { //обход дерева
		var nextlevel []*NodeOfTree //дети текущих узлов, для обработки на последующих итерациях

		for _, current := range currentlevel {
			result = append(result, current.data)

			if current.left != nil { //ищем левого ребенка
				nextlevel = append(nextlevel, current.left)
			}
			if current.right != nil { //ищем правого ребенка
				nextlevel = append(nextlevel, current.right)
			}
			currentlevel = nextlevel
		}
	}
	return result
}

// обход в глубину - рекурсивен, основан на стэке, левое поддерево(можно правое) -> вершина -> правое поддерево - симметричный обход, в результате отсортированная последовательность
func deptTraversal(root *NodeOfTree) []int {
	if root == nil {
		return []int{}
	}
	return deptTraversal(root.left)
	return deptTraversal(root.right)

	return []int{}
}

func dept(root *NodeOfTree, res int) int {
	if root == nil {
		return res
	}

	dept(root.left, res)
	res += root.data
	dept(root.right, res)

	return res
}

func DeptTraversalIterative(root *NodeOfTree) {
	stack := []*NodeOfTree{root} //кладём в стек корень дерева

	for len(stack) > 0 {
		current := stack[len(stack)-1] //верхний элемент стека
		stack = stack[:len(stack)-1]   //удаляем его
		fmt.Println(current.data)

		if current.left != nil {
			stack = append(stack, current.left)
		}
		if current.right != nil {
			stack = append(stack, current.right)
		}
	}
}
