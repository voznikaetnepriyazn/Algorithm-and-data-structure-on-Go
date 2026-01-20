package algs

//двусвязный список
type Node struct { //структура одного узла
	data int
	next *Node //ссылка на следующий элемент
	prev *Node
}
type LinkedList struct {
	head *Node //голова списка
	size int
	tail *Node //указатель на последний узел списка
}

//вставка в начало
func AppendFront(list *LinkedList, data int) {
	node := &Node{
		data: data,
		next: nil,
		prev: nil,
	}
	if list.head == nil {
		list.head = node
		return
	}

	//прежний head теперь в параметре next для нового узла
	node.next = list.head

	//новый узел становится предыдущим в параметр prev
	list.head.prev = node

	//записываем в head новый узел
	list.head = node
}

//вставка в конец
func AppendBack(list *LinkedList, data int) {
	node := &Node{
		data: data,
		next: nil,
		prev: nil,
	}
	if list.head == nil {
		list.head = node
		return
	}
	cur_node := list.head
	for cur_node.next != nil {
		cur_node = cur_node.next
	}

	//элементу, который был последним, в поле next записываем новый
	cur_node.next = node

	//в новый элемент в поле prev записываем узел, который до вставки был последним
	node.prev = cur_node
}

//развернуть список
func ReverseLinkedList(head *Node) *Node {
	prev := &Node{}
	current := head
	for current != nil {
		//сохраняем указатель на следующий узел во временную переменную
		next := current.next
		//next должен ссылаться на prev
		current.next = prev
		//двигаем указатель prev на единицу вперед - т е теперь current ссылается на него
		prev = current
		current = next
	}
	head = prev
	return head
}

//проверка является ли список циклическим - последний элемент ссылается на любой предыдущий
func HasCycle(head *Node) bool {
	if head == nil || head.next == nil {
		return false
	}
	slow := head
	fast := head.next
	for slow != fast {
		if fast == nil || fast.next == nil {
			return false
		}
		slow = slow.next
		fast = fast.next.next
	}
	return true
}
