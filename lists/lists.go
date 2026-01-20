package algs

//односвязный список - каждый элемент знает, где находится следующий элемент
type Node1 struct { //структура одного узла
	data int
	next *Node //ссылка на следующий элемент
}

type LinkedList1 struct {
	head *Node //голова списка
	size int
	tail *Node //указатель на последний узел списка
}

//вставка вначало списка - О(1)
func AddNewHead(list *LinkedList, el int) {
	node := &Node{
		data: el,
		next: nil,
	}

	if list.head == nil {
		list.tail = node
	}

	//прежний head сдвигаем на 1 узел назад
	//созданный узел в качестве next ссылается на head
	node.next = list.head

	list.head = node
}

//вставка в конец списка с указателем на tail - O(1)
func AddNewTail(list *LinkedList, el int) {
	node := &Node{
		data: el,
		next: nil,
	}

	//если список был пуст
	if list.tail == nil {
		list.head = node //одновременно и голова и хвост списка
	}
	list.tail.next = node //прежний tail ссылается в поле next на новый узел

	list.tail = node //новый узел в акчестве tail
}

//вставка в конец списка без указателя на tail - O(n)
func AddNewTail1(list *LinkedList, el int) {
	node := &Node{
		data: el,
		next: nil,
	}
	currentNode := list.head

	//нахождение последнего элемента
	for currentNode.next != nil {
		currentNode = currentNode.next
	}
	currentNode.next = node
}

//вставка в середину - O(n)
func Insert(list *LinkedList, after int, el int) *Node {
	search := list.head
	for search != nil {
		if search.data == after {
			break //нашли after
		}
		search = search.next //переходим к следующему элементу
	}

	node := &Node{
		data: el,
		next: nil,
	}

	//если нашли элемент after
	if search != nil {
		return node
	}

	//если элемент, после которого надо произвести вставку, является последним, то новый элемент теперь tail
	if search == list.tail {
		list.tail = node
	}

	//узел который был для search следующим, теперь next для нового узла
	node.next = search.next
	search.next = node

	return node
}

//поиск элемента - O(n)
func Search(head *Node, el int) *Node {
	current := head
	for current != nil {
		if current.data == el {
			return current
		}
		current = current.next
	}
	return nil
}

//найти середину списка
func MiddleNode(head *Node) *Node {
	var slow, fast *Node
	for fast != nil && fast.next != nil {
		slow = slow.next
		fast = fast.next.next
	}
	return slow
}
