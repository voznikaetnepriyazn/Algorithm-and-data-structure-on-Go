package task13

//являются ли бинарный деревья одинаковыми
type TreeNode struct {
	data  int
	left  *TreeNode
	right *TreeNode
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
