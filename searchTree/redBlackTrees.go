package algs

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// AVL ДЕРЕВО (через balanceFactor)
type NodeAVL struct {
	data          int
	parent        *NodeAVL
	left          *NodeAVL
	right         *NodeAVL
	balanceFactor int //разница высот: right - left
}

func NewNodeAVL(data int) *NodeAVL {
	return &NodeAVL{
		data:          data,
		parent:        nil,
		left:          nil,
		right:         nil,
		balanceFactor: 0,
	}
}

type AVLTree struct {
	root *NodeAVL
}

func NewAVLTree() *AVLTree {
	return &AVLTree{
		root: nil,
	}
}

// Малый левый поворот
func (t *AVLTree) leftRotate(x *NodeAVL) {
	y := x.right
	x.right = y.left

	if y.left != nil {
		y.left.parent = x
	}

	y.parent = x.parent

	if x.parent == nil {
		t.root = y
	} else if x == x.parent.left {
		x.parent.left = y
	} else {
		x.parent.right = y
	}

	y.left = x
	x.parent = y

	// Пересчёт баланс-факторов
	x.balanceFactor = x.balanceFactor - 1 - max(0, y.balanceFactor)
	y.balanceFactor = y.balanceFactor - 1 + min(0, x.balanceFactor)
}

// Малый правый поворот
func (t *AVLTree) rightRotate(y *NodeAVL) {
	x := y.left
	y.left = x.right

	if x.right != nil {
		x.right.parent = y
	}

	x.parent = y.parent

	if y.parent == nil {
		t.root = x
	} else if y == y.parent.right {
		y.parent.right = x
	} else {
		y.parent.left = x
	}

	x.right = y
	y.parent = x

	// Пересчёт баланс-факторов
	y.balanceFactor = y.balanceFactor + 1 - min(0, x.balanceFactor)
	x.balanceFactor = x.balanceFactor + 1 + max(0, y.balanceFactor)
}

// Вставка в AVL-дерево
func (t *AVLTree) Insert(key int) {
	node := NewNodeAVL(key)
	var parent *NodeAVL = nil
	cur := t.root

	// Поиск места для вставки
	for cur != nil {
		parent = cur
		if key < cur.data {
			cur = cur.left
		} else {
			cur = cur.right
		}
	}

	// Вставка узла
	node.parent = parent
	if parent == nil {
		t.root = node
	} else if key < parent.data {
		parent.left = node
	} else {
		parent.right = node
	}

	// Восстановление баланса
	t.updateBalance(node)
}

func (t *AVLTree) updateBalance(node *NodeAVL) {
	if node == nil {
		return
	}

	// Проверка на нарушение баланса
	if node.balanceFactor < -1 || node.balanceFactor > 1 {
		t.rebalance(node)
		return
	}

	// Обновление баланса родителя
	if node.parent != nil {
		if node == node.parent.left {
			node.parent.balanceFactor -= 1
		} else {
			node.parent.balanceFactor += 1
		}

		// Рекурсивное обновление вверх по дереву
		if node.parent.balanceFactor != 0 {
			t.updateBalance(node.parent) // ИСПРАВЛЕНО: node.parent вместо node
		}
	}
}

func (t *AVLTree) rebalance(node *NodeAVL) {
	if node.balanceFactor > 1 {
		// Правое поддерево тяжелее
		if node.right.balanceFactor < 0 {
			// Большое левое вращение (правый-левый)
			t.rightRotate(node.right)
			t.leftRotate(node)
		} else {
			// Малое левое вращение
			t.leftRotate(node)
		}
	} else if node.balanceFactor < -1 {
		// Левое поддерево тяжелее
		if node.left.balanceFactor > 0 {
			// Большое правое вращение (левый-правый)
			t.leftRotate(node.left)
			t.rightRotate(node)
		} else {
			// Малое правое вращение
			t.rightRotate(node)
		}
	}
}

// КРАСНО-ЧЕРНОЕ ДЕРЕВО (через цвета)
const (
	Red   = 1
	Black = 0
)

type NodeRB struct {
	data   int
	left   *NodeRB
	right  *NodeRB
	parent *NodeRB
	color  int
}

func NewNodeRB(data int) *NodeRB {
	return &NodeRB{
		data:   data,
		left:   nil,
		right:  nil,
		parent: nil,
		color:  Red, // Новый узел всегда красный
	}
}

type RedBlackTree struct {
	root *NodeRB
}

func NewRedBlackTree() *RedBlackTree {
	return &RedBlackTree{
		root: nil,
	}
}

// Левый поворот для красно-черного дерева
func (t *RedBlackTree) leftRotate(x *NodeRB) {
	y := x.right
	x.right = y.left

	if y.left != nil {
		y.left.parent = x
	}

	y.parent = x.parent

	if x.parent == nil {
		t.root = y
	} else if x == x.parent.left {
		x.parent.left = y
	} else {
		x.parent.right = y
	}

	y.left = x
	x.parent = y
}

// Правый поворот для красно-черного дерева
func (t *RedBlackTree) rightRotate(y *NodeRB) {
	x := y.left
	y.left = x.right

	if x.right != nil {
		x.right.parent = y
	}

	x.parent = y.parent

	if y.parent == nil {
		t.root = x
	} else if y == y.parent.right {
		y.parent.right = x
	} else {
		y.parent.left = x
	}

	x.right = y
	y.parent = x
}

// Вставка в красно-черное дерево
func (t *RedBlackTree) Insert(data int) {
	newNode := NewNodeRB(data)
	t.insertNode(newNode)
	t.fixInsert(newNode)
}

// Вспомогательный метод для вставки узла (без балансировки)
func (t *RedBlackTree) insertNode(node *NodeRB) {
	var parent *NodeRB = nil
	cur := t.root

	for cur != nil {
		parent = cur
		if node.data < cur.data {
			cur = cur.left
		} else {
			cur = cur.right
		}
	}

	node.parent = parent
	if parent == nil {
		t.root = node
	} else if node.data < parent.data {
		parent.left = node
	} else {
		parent.right = node
	}
}

// Восстановление свойств красно-черного дерева после вставки
func (t *RedBlackTree) fixInsert(node *NodeRB) {
	for node.parent != nil && node.parent.color == Red {
		if node.parent == node.parent.parent.left {
			// Родитель — левый потомок дедушки
			uncle := node.parent.parent.right

			if uncle != nil && uncle.color == Red {
				// Случай 1: дядя красный
				node.parent.color = Black
				uncle.color = Black
				node.parent.parent.color = Red
				node = node.parent.parent
			} else {
				// Случай 2: дядя черный
				if node == node.parent.right {
					node = node.parent
					t.leftRotate(node)
				}
				// Случай 3
				node.parent.color = Black
				node.parent.parent.color = Red
				t.rightRotate(node.parent.parent)
			}
		} else {
			// Симметричный случай: родитель — правый потомок дедушки
			uncle := node.parent.parent.left

			if uncle != nil && uncle.color == Red {
				// Случай 1: дядя красный
				node.parent.color = Black
				uncle.color = Black
				node.parent.parent.color = Red
				node = node.parent.parent
			} else {
				// Случай 2: дядя черный
				if node == node.parent.left {
					node = node.parent
					t.rightRotate(node)
				}
				// Случай 3
				node.parent.color = Black
				node.parent.parent.color = Red
				t.leftRotate(node.parent.parent)
			}
		}
	}

	// Корень всегда черный
	if t.root != nil {
		t.root.color = Black
	}
}
