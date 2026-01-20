package algs
import (
	"fmt"
	"strconv"
	"strings"
	"sort"
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

func GroupAnagrams(strs []string){
	anagrams := make(map[string][]string)

	for _, word := range strs{
		sorted_word := sort.Strings(strs)
	}
	if sorted_word, exists := anagrams[sorted_word]; exists{
		anagrams[sorted_word].PushBack(word)
	}
	anagrams[sorted_word] = [word]
	return anagrams
}