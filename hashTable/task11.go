package algs

import (
	"sort"
)

func FeedAnimals(animals, food []int) {
	if len(animals) == 0 || len(food) == 0 {
		return 0
	}
	sort.Ints(animals)
	sort.Ints(food)

	count := 0               //сколько можем накормить животных
	for _, f := range food { //сопоставляем текущую порцию с потребностью животного под индексом count
		if len(animals) > count {
			if f >= animals[count] {
				count += 1
			}
		}
		if count == len(animals) {
			break
		}
	}
	return count

}
