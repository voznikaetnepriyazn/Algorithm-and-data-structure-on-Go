package algs

//посчитать кол-во компонент связности
func FindConnectedComponents(graph map[int][]int) [][]int {
	visited := make(map[int]bool)  //для хранения посещенных вершин
	components := make([][]int, 0) //для хранения компонент связности

	//итерируемся по всем вершинам графа (ключам мапы)
	for vertex := range graph {
		if !visited[vertex] {
			var component []int
			Dfs24(graph, vertex, visited, &component)
			components = append(components, component)
		}
	}

	//обрабатываем изолированные вершины (которые есть в соседях, но не как ключи)
	for _, neighbors := range graph {
		for _, v := range neighbors {
			if !visited[v] {
				var component []int
				Dfs24(graph, v, visited, &component)
				components = append(components, component)
			}
		}
	}
	return components
}

func Dfs24(graph map[int][]int, v int, visited map[int]bool, component *[]int) { //v - вершина
	visited[v] = true
	*component = append(*component, v)

	for _, neighbor := range graph[v] {
		if !visited[neighbor] {
			Dfs24(graph, neighbor, visited, component)
		}
	}
}
