package algs

//поиск минимальной глубины бинарного дерева - количество узлов на кратчайшем пути от корневого узла до близжайшего листового узла
type TreeNode struct {
	data  int
	left  *Node
	right *Node
}

func minDepth(root *TreeNode) *TreeNode {
	if root == nil {
		return 0
	}

	if root.left == nil && root.right == nil {
		return 1 + min(minDepth(root.left), minDepth(root.right)) //единицу прибавляем чтобы учесть корень
	}
	if root.left != nil {
		return 1 + minDepth(root.left)
	}
	if root.right != nil {
		return 1 + minDepth(root.right)
	}
}
