package algs

//поиск цикла в графе
//граф задан в виде словаря вершин
//если в процессе dfs соседняя вершина уже была посещена и не являетс родительской для текущей вершины, в графе есть цикл

func Dfs(graph map[string][]string, vertex int, parent string, visited map[string]bool) bool {
	visited[vertex] = true

	for _, neighbor := range graph[vertex] {
		if neighbor != parent {
			if visited[neighbor] || Dfs(graph, neighbor, vertex, visited) {
				return true
			}
		}
	}
	return false
}

func HasCycle(graph map[string][]string) {
	visited := make(map[string]bool) //массив посещенных узлов

	for vertex := range graph {
		if !visited[vertex] {
			if Dfs(graph, vertex, "", visited) {
				return true
			}
		}
	}
	return false
}
