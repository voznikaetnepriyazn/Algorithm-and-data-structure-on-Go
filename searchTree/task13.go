package algs

//являются ли бинарный деревья одинаковыми
type TreeNode struct {
	data  int
	left  *Node
	right *Node
}

func isSameTree(a, b *TreeNode) any {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	if a.data != b.data {
		return false
	}
	return isSameTree(a.left, b.left) && isSameTree(a.right, b.right)

}
