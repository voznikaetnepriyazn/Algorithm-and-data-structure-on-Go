package algs

//O(1)
type Deque struct{
	head,tail *Node
	size int
}

type Node struct{
	value int
	prev, next *Node
}

func NewDeque() *Deque{
	head := &Node{}
	tail := &Node{}
	head.next := tail
	tail.prev := head
	return &Deque{
		head : head,
		tail : tail,
		size : 0,
	}
}

func (d *Deque) PushFront(value int){
	newNode := &Node{
		value := value
		prev = d.head
		next = d.head.next
	}
	newNode.next := d.head.next//ссылка на некогда 1ый элемент
	newNode.prev := d.head//ссылка на предыдующий на головной элемент
	d.head.next.prev := newNode//некогда первый элемент настроить на новый
	d.head.next := newNode//голову настроить на новый
	d.size++
}

func (d *Deque) PopBack(){
	if d.head.next == d.tail{
		return null
	}
	popResult := d.tail.prev//извлекаем всегда из начала очереди
	d.tail.prev = popResult.prev//теперь хвост в качестве prev(1го элемента очереди) ссылается на 2ой элемент
	popResult.prev.next = popResult.next//теперь next ссылается на tail
	//отцепляем элемент от списка
	popResult.next = nil
	popResult.prev = nil
	d.size--
	return popResult.data
}

func (d *Deque) PushBack(value int){
	newNode := &Node{
		value := value
		prev = d.head
		next = d.head.next
	}
	newNode.next := d.tail
	newNode.prev := d.tail.prev
	d.tail.prev.next = newNode
	d.tail.prev = newNode
	d.size++
}

func (d *Deque) PopFront(){
	if d.head.next == d.tail{
		return null
	}
	popResult := d.head.next
	d.head.next = popResult.next
	popResult.next.prev = popResult.prev
	popResult.next = nil
	popResult.prev = nil
	d.size--
	return popResult.datapopResult.next = nil
}