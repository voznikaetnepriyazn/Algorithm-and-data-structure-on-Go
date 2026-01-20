package algs
import (
	"fmt"
	"strconv"
	"strings"
	"math"
)
//восстановление бинарного дерева из массива
type TreeNode struct{
	data int
	left *TreeNode
	right *TreeNode
}
func TreeNode(data int, left *Node, right *Node) *TreeNode{
	return &TreeNode{
		data := 0
		left := nil
		right := nil
	}
}

func buildTree(arr []int, i int) *TreeNode{//i - индекс, для которого мы должны создать поддерево
	if i > len(arr){
		return nil
	}
	root := &TreeNode{ Val : arr[i]}
	root.left = buildTree(arr , 2*i+1)
	root.right = buildTree(arr , 2*i+2)

	return root
}