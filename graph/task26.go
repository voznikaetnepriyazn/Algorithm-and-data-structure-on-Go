package algs

import (
	"container/heap"
	"math"
)

//алгоритм дейкстры - поиск кратчайшего пути
//взвешенный граф - у ребер есть занчения
//для вершины ставим значение 0

type Item struct { //элемент в очереди с приоритетом
	vertex   string
	distance float64
	index    int
}

type PriorityQueue []*Item

func (q *PriorityQueue) Len() int {
	return len(q)
}

func (q *PriorityQueue) Less(i, j int) bool {
	return q[i].distance < q[j].distance
}

func (q *PriorityQueue) Swap(i, j int) {
	q[i], q[j] = q[j], q[i]
	q[i].index = i
	q[j].index = j
}

func (q *PriorityQueue) Push(x any) {
	n := len(*q)
	item := x.(*Item)
	item.index = n
	*q = append(*q, item)

}

func (q *PriorityQueue) Pop() any {
	old := *q
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.index = -1
	*q = old[0 : n-1]
	return item
}

func Dijkstra(graph map[string]map[string]float64, start string) {
	distances := make(map[string]float64)
	for vertex := range graph {
		distances[vertex] = math.Inf(1)
	}
	distances[start] = 0

	q := *&PriorityQueue{}
	heap.Init(q)
	heap.Push(q, &Item{vertex: start, distance: 0})

	for q.Len > 0 {
		current := heap.Pop(q).(*Item)
		currentVertex := current.vertex
		currentDistance := current.distance

		if currentDistance > distances[currentVertex] {
			continue
		}

		for neighbor, weight := range graph[currentVertex] {
			distance := currentDistance + weight
			if distance < distances[neighbor] {
				distances[neighbor] = distance
				heap.Push(q, &Item{vertex: neigbor, distance: distance})
			}
		}
	}
	return distances
}
