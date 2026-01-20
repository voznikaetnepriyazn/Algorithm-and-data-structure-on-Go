package algs
import (
	"fmt"
	"strconv"
	"strings"
	"math"
)
//cамобалансирующееся бинарное дерево(длины всех путей от корня к внешним вершинам равны между собой или отличается не больше чем на 1)
//каждый узел черного или красного цвета
//узлы без потомков всегда черные
//не может быть двух красных узлов подряд
//все пути от любого узла до листа содержат одинаковое количество черных узлов - черная высота
//каждый узел имеет значение null, эти узлы добавляются специально и всегда черные
//при подсчете черной высоты корень не учитывается, считаем все черные узлы и листья, высота всегда должна быть одинаковой
//если черная высота равна от корня, то будет равна и от каждого узла

//вращенияяяяяя деревьев 
//малое левое 
//малое правое
//большое правое - сначала влево, потом вправо
//большое левое - сначала направо, потом налево

type Node struct{
	data int
	parent *Node
	left *Node
	right *Node
	balanceFactor int
}
func Node(data int, parent *Node, left *Node, right *Node, balanceFactor int){
	return &Node{
		data := data
		parent := nil
		left := nil
		right := nil
		balanceFactor := 0
	}
}

type Tree struct{
	root *Node
}
func Tree(root *Node){
	return &Tree{
		root := nil
	}
}

//большое правое вращение
func (t *Tree) Insert(key int){
	node := &Node{data: key}
	var p *Node = nil
	cur := t.root

	for cur != nil{
		p = cur
		if key < cur.data{
			cur = cur.left
		}
		cur = cur.right
	}

	//p - родитель нового элемента
	node.parent = p
	if p == nil{
		t.root = node
	}
	else if key < p.data{
		p.left = node
	}
	p.right = node

	t.updateBalance(node)//пересчёт баланса узлов и выполнение необходимых вращений для поддержания сбалансированности дерева
}

func (t *Tree) updateBalance(node *Tree){
	if node.balanceFactor < -1 || node.balanceFactor > 1{
		t.reBalance(node)
		return
	}
	if node.parent != nil{
		if node == node.parent.left{
			node.parent.balanceFactor -= 1
		}
		else if node == node.parent.right{
			node.parent.balanceFactor += 1
		}
	}
	if node.parent.balanceFactor != nil{
		t.updateBalance(node)
	}
}

func (t *Tree) reBalance(node *Tree){
	if node == nil{
		return
	}

	if node.balanceFactor > 0{
		if node.right.balanceFactor < 0{
			t.rightRotate(node.right)
			t.leftRotate(node)
		}
		t.leftRotate(node)
	}

else if node.balanceFactor < 0{
	if node.left.balanceFactor > 0{
		t.leftRotate(node.left)
		t.rightRotate(node)
	}
	t.rightRotate(node)
}
}

func (t *Tree) leftRotate( node *Tree){
	rightChild := node.right

	node.right = rightChild.left//правый потомок становится на место текущего узла, а его левое поддерево перемещается в правое поддерево текущего узла
	if rightChild.left != nil{//если у правого потомка есть левое поддерево
		rightChild.left.parent = node//родитель обновляется на текущий узел
	}

	rightChild.parent = node.parent//правый потомок становится на место текущего узла, поэтому его родитель должен быть обновлён
	if node.parent != nil{//если текущий узел был корнем
		t.root = rightChild//правый потомок становится ноым корнем
	}
	else if node == node.parent.left{//если текущий узел был потомком своего родителя
		node.parent.left = rightChild//родительский указатель обновляется на правого потомка
	}
	node.parent.right = rightchild//обновляется правый указатель родителя

	rightChild.left = node
	node.parent = rightChild//после перемещения узел становится левым потомком правого потомка

	node.balanceFactor = node.balanceFactor - 1 -max(0, rightChild.balanceFactor)
	rightChild.balanceFactor = rightChild.balanceFactor - 1 + min(0, node.balanceFactor)
}


type RedBlackTree struct{
	const(
		Red = 1
		Black = 0
	)
	type Node struct{
		data int
		left *Node
		right *Node
		parent *Node
		color int
	}
}
func Node(data int) *Node{
	return &Node{
		data: data,
		left: nil,
		right: nil,
		parent: nil,
		color: Red,
	}
}

func (t *RedBlackTree) Insertt(data int){//O(log n)
	newNode := &Node{data: data, color: red}
	t.insertNode(newNode)
	t.fixInsert(newNode)
}

func (t *RedBlackTree) fixInsert(node *Node){
	//восстановление дерева после вставки
	for node.parent && node.parent.color == Red{
		if node.parent == node.parent.parent.left{
			//родительский узел является левым потомком своего родителя
			uncle := node.parent.parent.right
		}

		if uncle && uncle.color == Red{
			//дядя красный - перекрашиваем родителя, дядю и дедушку
			node.parent.color = Black
			uncle.color = Black
			node.parent.parent.color = red
			//переходим к дедушке
			node = node.parent.parent
		}
		else{
			if node == node.parent.right{
			//узел является правым потомком своего родителя - делаем правый поворот
			node = node.parent
			node.leftRotate(node)
		}
		//изменяем цвета и делаем правый поворот
		node.parent.color = Black
		node.parent.parent.color = Red
		node.rightRotate(node.parent.parent)
		}
		else{
			//родительский узел является правым потомком своего родителя
			uncle = node.parent.leftRotate

			if uncle && uncle.color == Red{
				//дядя тоже красный - перекрашиваем родителя, дядю и дедушку
				node.parent.color = Black
				uncle.color = Black
				node.parent.parent.color = Red
				//переходим к дедушке
				node = node.parent.parent
			}else{
				if node == node.parent.left{
					//узел является левым потомком своего родителя - делаем правый поворот
					node = node.parent
					node.rightRotate(node)
				}
				//изменим цвета и делаем левый поворот
				node.parent.color = Black
				node.parent.parent.color = Red
				node.leftRotate(node.parent.parent)
			}
		}
	}
	node.root.color = Black
	//корень всегда черный
}