package algs

//двусвязный список
type Node2 struct { //структура одного узла
	data int
	next *Node //ссылка на следующий элемент
	prev *Node
}
type LinkedList2 struct {
	head *Node //голова списка
	size int
	tail *Node //указатель на последний узел списка
}

func SuumOf(n int, arr []int) (int, int) {
	current := &Node{}
	min := current.next.data - current.data
	first, second := current.data, current.next.data

	if n < 2 {
		return 0, 0
	}

	for i := 1; i < n-1; i++ {
		if current.next.data-current.data < min {
			min = current.next.data - current.data
			first, second = current.data, current.next.data
		}
		current = current.next
	}
	return first, second
}
