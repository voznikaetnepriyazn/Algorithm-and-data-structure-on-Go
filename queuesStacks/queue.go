package algs
import (
	"fmt"
	"strconv"
	"strings"
)
//first in first out O(1)

type Queue struct {
	data []int{}
	head, tail int
	size int
	capacity int
}

func NewQueue(cap int) *Queue{
	if cap <= 0{
		cap = 4
	}
	return &Queue{
		data := make([]int{}, cap)
		capacity := cap
	}
}

func (a *Queue) Enqueue(value int) {  //О(1)
	if a.size == a.capacity{
		a.resize()
	}
	a.data[a.tail] = value//при добавлении устанавливать элемент на индекс tail
	a.tail = (a.tail + 1) % a.capacity//двигает хвост на единицу к концу массива, деление на capacity - реализация кольцевого буфера
	a.size++
}

func (b *Queue) Dequeue(value int) {
	if b.size == 0{ 
		return 0
	}
	value := b.data[b.head]
	b.head = (b.head + 1) % b.capacity
	b.size--
	return value, nil
}

func (q *Queue) Resize() {
	newCap := q.capacity * 2
	newData := make([]int{}, newcap)
	for i := 0, i < q.size, i++{//копирование всех элементов в новый массив
		newData[i] = q.data[(q.head + i) % q.capacity]
	}
	q.data = newData
	q.head = 0
	q.tail = q.size
	q.capacity = newCap
}

