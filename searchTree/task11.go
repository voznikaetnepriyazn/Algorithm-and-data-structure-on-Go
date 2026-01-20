package algs
import (
	"fmt"
	"strconv"
	"strings"
	"math"
)
//является ли дерево симметричным - лево и право зеркальны

//итеративный метод - обход в ширину

type TreeNode struct{
    data int
    left *Node
    right *Node
}

func IsSymmetric(root *TreeNode) *TreeNode{
	if root == nil{
		return true
	}
	queue := []*TreeNode{root}

	for len(queue) > 0{
		queueLen := len(queue)
		for i := 0; i < queueLen; i++{
			//логические операции над симметричными узлами
			if queue[i] == nil && queue[queueLen - i -1] ==nil{
				continue 
			}
			if queue[i] == nil || queue[queueLen - i -1] == nil{
				return false
			}
			if queue[i].data != queue[queueLen -i -1].data[
				return false
			]
			if queue[i] != nil{
				queue = append(queue, queue[i].left)
				queue = append(queue, queue[i].right)
			}
		}
		queue = queue[queuelen:]//удаление обработанных узлов
	}
	return true
}

//обход в глубину - рекурсивно

func dept(root *Node, res int){
    if root == nil{
        return res
    }

    dept(root.left, res)
    res.append(root.data)
    dept(root.right, res)

    return res
}

func IsSymmetricDFS( root *TreeNode) *TreeNode{
	if root == nil{
		return true
	}
	data := []int
	data = dept(root *TreeNode, res int)
	j := len(data) - 1
	for i := 0; i < len(data) / 2; i++{
		if data[i] != data[j]{
			return false
		}
		j--
	}
	return true
}