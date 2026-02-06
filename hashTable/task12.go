package algs

import (
	"sort"
)

type Deque struct {
	head, tail *Node
	size       int
}

type Node struct {
	value      int
	prev, next *Node
}

func NewDeque() *Deque {
	head := &Node{}
	tail := &Node{}
	head.next = tail
	tail.prev = head
	return &Deque{
		head: head,
		tail: tail,
		size: 0,
	}
}

func (d *Deque) PushBack(value int) {
	newNode := &Node{
		value: value,
		prev:  d.head,
		next:  d.head.next,
	}
	d.tail.prev.next = newNode
	d.tail.prev = newNode
	d.size++
}

func GroupAnagrams(strs []string) [][]string {
	anagrams := make(map[string][]string)

	for _, word := range strs {
		chars := []rune(word)
		sort.Slice(chars, func(i, j int) bool { return chars[i] < chars[j] })
		key := string(chars)
		anagrams[key] = append(anagrams[key], word)
	}

	result := make([][]string, 0, len(anagrams))
	for _, group := range anagrams {
		result = append(result, group)
	}
	return result
}
