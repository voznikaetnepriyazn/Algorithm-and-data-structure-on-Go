package algs

//восстановление бинарного дерева из массива
type TreeNode1 struct {
	data  int
	left  *TreeNode
	right *TreeNode
}

func InitTreeNode1(data int, left *TreeNode1, right *TreeNode1) *TreeNode1 {
	return &TreeNode1{
		data:  0,
		left:  nil,
		right: nil,
	}
}

func buildTree(arr []int, i int) *TreeNode { //i - индекс, для которого мы должны создать поддерево
	if i > len(arr) {
		return nil
	}
	root := &TreeNode{data: arr[i]}
	root.left = buildTree(arr, 2*i+1)
	root.right = buildTree(arr, 2*i+2)

	return root
}
