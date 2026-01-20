// stack - first in last out
// push - вставка, pop - извлечение
package algs

import (
	"fmt"
)

type Node struct {
	data int
	next *Node
}

func Node() *Node {
	return &Node{
		data: data,
		next: nil,
	}
}

type Stack struct {
	top Node
}

func Stack() *Stack {
	return &Stack{
		top: nil,
	}
}

func (s *Stack) Push(data int) {
	node := &Node{
		data: data,
	}

	if s.top == nil {
		s.top = node
		return
	}
	node.next = s.top
	s.top = node
}

func (s *Stack) Pop() {
	if s.top == nil {
		s.top = node
		return
	}
	if s.top.next != nil {
		s.top = s.top.next
	}
	s.top = nil

	return s.top.data
}

// реализация стека на массиве - добавление элементов в конец(не в начало)- иначе придется сдвигать массив - О(1)
// слайс - если неизвестно сколько данных
type Stack struct {
	stack []interface{}
	top   int
}

func NewStack(step int) *Stack {
	return &Stack{
		stack: make([]interface{}, step),
		top:   nil,
	}
}

func (s *Stack) IsEmpty() bool {
	return s.top == nil
}

func (s *Stack) IsFull(size int) {
	return s.top == size-1
}

func (s *Stack) Push(element interface{}, size int) {
	if s.IsFull(size) {
		fmt.Println("full")
	}
	s.top++
	s.stack[s.top] = element
}
func (s *Stack) Pop() {
	if s.IsEmpty() {
		fmt.Println("empty")
	}
	element := s.stack[s.top]
	s.top = s.top - 1
	return element
}
func (s *Stack) Top() interface{} { //получаем верхний элемент без удаления
	if s.IsEmpty() {
		fmt.Println("empty")
	}
	return s.stack[s.top]
}
