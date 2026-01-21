package algs

//куча - бинарное дерево, только все значения ниже вершины меньше вершины( и лево и право)
//левый потомок - 2*i+1, правый - 2*i+2, индексы
//родительский узел - (i - 1)/2

type MaxHeap struct {
	data []int
}

//создание кучи из массива
func buildMaxHeap(data []int) { //O(n*log n)
	n := len(data)
	//начинаем с последнего уровня, имеющего потомков
	for i := n/2 - 1; i >= 0; i-- {
		siftDown(data, n, i)
	}
}

func siftDown(data []int, n, i int) { //восстановление свойств бинарной кучи после изменений элемента
	largest := i
	left := 2*i + 1
	right := 2*i + 2

	if left < n && data[left] > data[largest] {
		largest = left
	}
	if right < n && data[right] > data[largest] {
		largest = right
	}
	if largest != i {
		//опускаем вниз узел с меньшим значением
		data[i], data[largest] = data[largest], data[i] //а больший узел поднимаем вверх
		siftDown(data, n, largest)
	}
}

//очередь с приоритетом
//О(log n)- получение значения pop
//О(1) - доступ к максимальному значению

func top() {
	return heap[0]
}
func right() {
	return 2*i + 2
}
func left() {
	return 2*i + 1
}
func parent(i int) {
	if i == 0 {
		return i
	}
	return (i - 1) / 2
}

//добавление элемента
func (h *MaxHeap) push(data int) {
	h.heap = append(h.heap, data)
	h.up(len(h.data) - 1)
}

func (h *MaxHeap) up(index int) {
	//в цикле, пока добавляемый элемент больше родительского значения,меняем их местами и смещаем индекс на индекс ролительского узла
	for index > 0 {
		parent := (index - 1) / 2
		if h.data[index] <= h.data[parent] {
			h.data[index], h.data[parent] = h.data[parent], h.data[index]
			index = parent //смещение индекса
		}
	}
}

//удаление элемента на вершине кучи
func (h *MaxHeap) DeleteMax(data []int) int {
	if len(h.data) == 0 {
		return 0
	}
	//на место удалённого узла передвигаем узел из самого нижнего ряда справа и опускаем его вниз, меняя местами с максимальным потомком
	n := len(h.data)        //вершина кучи
	result := h.data[0]     //
	h.data[0] = h.data[n-1] // заменяем корневой элемент на последний элемент в куче
	h.data = h.data[:n-1]
	siftDown(data, n-1, 0)

	return result
}

//heap-sort O(n * log n)
type Minheap struct {
	data []int
}

func (h *Minheap) HeapSort(data []int) []int {
	n := len(h.data)
	for i := n/2 - 1; i >= 0; i-- {
		siftDown(data, n, i)
	} //построение максимальной кучи

	//извлечение элементов по одному и сортировка
	for i := n - 1; i > 0; i-- {
		//отправляем нулевой элемент в конец неотсортированной части массива, т к это теперь самый большой элемент
		data[0], data[i] = data[i], data[0]
		siftDown(data, i, 0)
	}
	return data
}
