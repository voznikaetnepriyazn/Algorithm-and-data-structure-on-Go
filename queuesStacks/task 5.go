package algs

//Deque (двусвязная очередь)
type Node3 struct {
	value int
	prev  *Node3
	next  *Node3
}

type Deque struct {
	head, tail *Node3
	size       int
}

func NewDeque() *Deque {
	head := &Node3{}
	tail := &Node3{}
	head.next = tail
	tail.prev = head
	return &Deque{
		head: head,
		tail: tail,
		size: 0,
	}
}

func (d *Deque) PushBack(value int) {
	newNode := &Node3{
		value: value,
		prev:  d.tail.prev,
		next:  d.tail,
	}
	d.tail.prev.next = newNode
	d.tail.prev = newNode
	d.size++
}

func (d *Deque) PopFront() int {
	if d.size == 0 {
		panic("deque is empty")
	}
	node := d.head.next
	d.head.next = node.next
	node.next.prev = d.head
	d.size--
	return node.value
}

func (d *Deque) PopBack() int {
	if d.size == 0 {
		panic("deque is empty")
	}
	node := d.tail.prev
	d.tail.prev = node.prev
	node.prev.next = d.tail
	d.size--
	return node.value
}

func (d *Deque) Len() int {
	return d.size
}

//Stack (стек на односвязном списке)
type Node4 struct {
	data int
	next *Node4
}

func InitNode4(data int) *Node4 {
	return &Node4{
		data: data,
		next: nil,
	}
}

type Stack3 struct {
	top  *Node4
	size int
}

func NewStack() *Stack3 {
	return &Stack3{
		top:  nil,
		size: 0,
	}
}

func (s *Stack3) Push(data int) {
	node := InitNode4(data)
	node.next = s.top
	s.top = node
	s.size++
}

func (s *Stack3) IsEmpty() bool {
	return s.top == nil
}

func (s *Stack3) Pop() int {
	if s.IsEmpty() {
		panic("stack is empty")
	}
	value := s.top.data
	s.top = s.top.next
	s.size--
	return value
}

func (s *Stack3) Size() int {
	return s.size
}

//Проверка подпоследовательности
func IsSubsequence(a, b string) bool {
	i := 0 // указатель для строки a
	for j := 0; j < len(b) && i < len(a); j++ {
		if a[i] == b[j] {
			i++
		}
	}
	return i == len(a)
}

// Альтернативная реализация (та же логика)
func IsSubsequenceVar(a, b string) bool {
	i, j := 0, 0
	for i < len(a) && j < len(b) {
		if a[i] == b[j] {
			i++
		}
		j++
	}
	return i == len(a)
}

//Проверка палиндрома
// Через стек (используем слайс как стек)
func IsPalindrome(s string) bool {
	stack := []rune{}
	for _, ch := range s {
		stack = append(stack, ch)
	}
	for _, ch := range s {
		if ch != stack[len(stack)-1] {
			return false
		}
		stack = stack[:len(stack)-1]
	}
	return true
}

// Через дек
func IsPalindromeVar(s string) bool {
	deq := NewDeque()
	for _, ch := range s {
		deq.PushBack(int(ch))
	}
	for deq.Len() > 1 {
		if deq.PopFront() != deq.PopBack() {
			return false
		}
	}
	return true
}

// Через два указателя (самый эффективный способ)
func IsPalind(s string) bool {
	left, right := 0, len(s)-1
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}

//Удаление элемента из односвязного списка
type ListNode1 struct {
	data int
	next *ListNode1
}

func NewListNode(data int) *ListNode1 {
	return &ListNode1{
		data: data,
		next: nil,
	}
}

// Исправлена логика удаления: всегда двигаем current вперед
func RemoveElement(head *ListNode1, val int) *ListNode1 {
	dummy := &ListNode1{next: head}
	prev := dummy
	current := head

	for current != nil {
		if current.data == val {
			prev.next = current.next // пропускаем текущий узел
		} else {
			prev = current // двигаем prev только если не удалили
		}
		current = current.next // всегда двигаем current
	}
	return dummy.next
}

//Поиск последнего четного числа
func Found(s []int) int {
	last := -1
	for i := len(s) - 1; i >= 0; i-- {
		if s[i]%2 == 0 {
			last = s[i]
			break
		}
	}
	return last
}
