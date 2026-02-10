// stack - first in last out
// push - вставка, pop - извлечение
package algs

import (
	"fmt"
)

type Node1 struct {
	data int
	next *Node1
}

func InitNode1(data int) *Node1 {
	return &Node1{
		data: data,
		next: nil,
	}
}

type Stack1 struct {
	top *Node1
}

func InitStack() *Stack1 {
	return &Stack1{
		top: nil,
	}
}

func (s *Stack1) Push(data int) {
	node := &Node1{
		data: data,
	}

	if s.top == nil {
		s.top = node
		return
	}
	s.top = node
}

func (s *Stack1) Pop() (int, error) {
	if s.top == nil {
		return 0, fmt.Errorf("стек пуст")
	}
	data := s.top.data
	s.top = s.top.next
	return data, nil
}

// реализация стека на массиве - добавление элементов в конец(не в начало)- иначе придется сдвигать массив - О(1)
// слайс - если неизвестно сколько данных
type Stack2 struct {
	stack []interface{}
	top   int
}

func NewStack2(step int) *Stack2 {
	return &Stack2{
		stack: make([]interface{}, step),
	}
}

func (s *Stack2) IsEmpty() bool {
	return s.top == 0
}

func (s *Stack2) IsFull(size int) bool {
	return s.top == size-1
}

func (s *Stack2) Push(element interface{}, size int) {
	if s.IsFull(size) {
		fmt.Println("full")
	}
	s.top++
	s.stack[s.top] = element
}
func (s *Stack2) Pop() interface{} {
	if s.IsEmpty() {
		fmt.Println("empty")
	}
	element := s.stack[s.top]
	s.top = s.top - 1
	return element
}
func (s *Stack2) Top() interface{} { //получаем верхний элемент без удаления
	if s.IsEmpty() {
		fmt.Println("empty")
	}
	return s.stack[s.top]
}
