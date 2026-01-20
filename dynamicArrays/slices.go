package algs

import "fmt"

//реаллокация - увеличение capacity

type DynamicArray struct {
	count    int
	size     int
	capacity int
	arr      []interface{}
}

//конструктор
func InitDynamicArray() *DynamicArray {
	capacity := 1
	return &DynamicArray{
		count:    1,
		size:     0,
		capacity: capacity,
		arr:      make([]interface{}, capacity),
	}
}

func (da *DynamicArray) Apppend(element int) {
	if da.size == da.capacity {
		da.Resize()
		fmt.Printf("slow: %d\n", da.capacity)
	}
	fmt.Printf("fast")
	da.arr[da.size] = element
	da.size++
}

func (da *DynamicArray) Resize() {
	newCap := da.capacity * 2
	newArr := make([]interface{}, 0, newCap)
	for i := 0; i < da.size; i++ {
		da.count++
		newArr[i] = da.arr[i]
	}
	da.arr = newArr
	da.capacity = newCap
}
