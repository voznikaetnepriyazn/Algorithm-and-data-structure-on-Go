package algs

//поиск минимальной глубины бинарного дерева - количество узлов на кратчайшем пути от корневого узла до близжайшего листового узла
type TreeNode struct {
	data  int
	left  *TreeNode
	right *TreeNode
}

// min возвращает минимальное из двух целых чисел
func minn(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// minDepth возвращает минимальную глубину дерева (количество узлов от корня до ближайшего листа)
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	// Если оба ребенка существуют — ищем минимум из двух поддеревьев
	if root.left != nil && root.right != nil {
		return 1 + minn(minDepth(root.left), minDepth(root.right))
	}

	// Если только левый ребенок существует
	if root.left != nil {
		return 1 + minDepth(root.left)
	}

	// Если только правый ребенок существует
	if root.right != nil {
		return 1 + minDepth(root.right)
	}

	// Если это лист (оба ребенка nil)
	return 1
}
