package algs
import (
	"fmt"
	"strconv"
	"strings"
)
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

type Node struct{
	data int
	next *Node
}

func Node() *Node{
	return &Node{
		data := data
		next := nil
	}
}

type Stack struct{
	top Node
}

func Stack() *Stack{
	return &Stack{
		top : nil
	}
}

//является ли одна строка исходной для другой

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
func (s *Stack) Push(data int){
	node := &Node{
		data : data
	}

	if s.top == nil{
		s.top = node
		return
	}
	node.next = s.top
	s.top = node
}
func (s *Stack) Pop() {
	if s.IsEmpty(){
		fmt.Println("empty")
	}
	element := s.stack[s.top]
	s.top = s.top - 1
	return element
}

func IsSubsequence(a,b string) bool{
	q := q.list.New()

	for _, ch := range a{
		q.PushBack(ch)
	}

	for _, ch := range b{
		if q.Len() == 0{
			break
		}
		if q.Front().value == ch{
			q.Pop(q.Front())
		}
	}
	return q.Len() == 0
	
}

func IsSubsequenceVar(a, b string){  //O(n)
	i, j := 0, 0
	for i < len(a) && j < len(b){
		if a[i] == b[j]{
			i++
		}
		j++
	}
	return i == len(a)
}

//является ли слово палиндромом  O(n)
func IsPalindrome(s string){
	stack := make([]interface{}, len(s))//используем стэк
	for _, i range s{
		stack.Push(i)
	}
	for _, j range s{
		if j != stack.Pop(){
			return false
		}
	}
	return true
}
func IsPalindromeVar(s string){//используем дэк
	var deq Deque
	for _, i range deq{
		deq.PushBack(i)
	}
	for deq.Len() > 1{
		if deq.head != deq.tail{
			return false
		}
		deq.PopFront(deq.head)
		deq.PopBack(deq.tail)
	}
	return true
}

func IsPalind(s string){//метод 2ух указателей O(n)
	left := 0
	right := len(s) - 1
	for left < right{
		if s[left] != s[right]{
			return false
		}
		left++
		right--
	}
	return true
}
//удалить элемент из односвязного списка - добавить фиктивный узел перед head
type ListNode struct{
	data int
	next *ListNode
}

func ListNode *ListNode{
	return &ListNode{
		data := data
		next := nil
	}
}

func RemoveElement(head *ListNode, val int){  //O(n)
	var dummy ListNode
	dummy.next := head
	prev := dummy
	current = head

	for current != nil{
		if current.data == val{
			prev.next = current.next//двигаем на узел вперед
		}
		prev = current//у предыдущего узла перепишем значение в поле некст
		current = current.next//двигаем на узел вперед
	}
	return dummy.next
}

func Found(s []interface{}, n int){
	arr := make([]interface{}, n)
	for _, i range s{
		arr.PushBack(i)
	}
	last := -1
	for i := n, i >= 0, i--{
		if arr[i] % 2 == 0{
			last := arr[i]
			break
		}
	}
	return last
}