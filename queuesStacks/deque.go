package algs

//O(1)
type Dequee struct {
	head, tail *Nodee
	size       int
}

type Nodee struct {
	value      int
	prev, next *Nodee
}

func InitDequee() *Dequee {
	head := &Nodee{}
	tail := &Nodee{}
	head.next = tail
	tail.prev = head
	return &Dequee{
		head: head,
		tail: tail,
		size: 0,
	}
}

func (d *Dequee) PushFront(value int) {
	newNode := &Nodee{
		value: value,
		prev:  d.head,
		next:  d.head.next,
	}
	newNode.next = d.head.next //ссылка на некогда 1ый элемент
	newNode.prev = d.head      //ссылка на предыдующий на головной элемент
	d.head.next.prev = newNode //некогда первый элемент настроить на новый
	d.head.next = newNode      //голову настроить на новый
	d.size++
}

func (d *Dequee) PopBack() int {
	if d.head.next == d.tail {
		return 0
	}
	popResult := d.tail.prev             //извлекаем всегда из начала очереди
	d.tail.prev = popResult.prev         //теперь хвост в качестве prev(1го элемента очереди) ссылается на 2ой элемент
	popResult.prev.next = popResult.next //теперь next ссылается на tail
	//отцепляем элемент от списка
	popResult.next = nil
	popResult.prev = nil
	d.size--
	return popResult.value
}

func (d *Dequee) PushBack(value int) {
	newNode := &Nodee{
		value: value,
		prev:  d.head,
		next:  d.head.next,
	}
	newNode.next = d.tail
	newNode.prev = d.tail.prev
	d.tail.prev.next = newNode
	d.tail.prev = newNode
	d.size++
}

func (d *Dequee) PopFront() bool {
	if d.head.next == d.tail {
		return false
	}
	popResult := d.head.next
	d.head.next = popResult.next
	popResult.next.prev = popResult.prev
	popResult.next = nil
	popResult.prev = nil
	d.size--
	return popResult.next == nil
}
