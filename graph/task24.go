package algs

//посчитать кол-во компонент связности
func FindConnectedComponents(graph map[int][]int) [][]int {
	visited := make(map[int]bool)        //для хранения посещенных вершин
	connectedComponents := make([][]int) //для хранения компонент связности

	for i := 1; i <= len(graph); i++ {
		visited[i] = false
	}

	//в цикле по графу для каждой непосещенной вершины запускаем dfs
	for i := 1; i <= len(graph); i++ {
		currentNode := graph[i]
		if !visited[currentNode] {
			//для каждого запуска создаем массив, в котором храним текущий подграф
			component := make([]int)
			Dfs(graph, currentNode, visited, component)
			//добавляем массив в connectedComponents
			connectedComponents = append(connectedComponents, component)
		}
	}
	return connectedComponents
}

func Dfs(graph map[string][]string, v int, visited map[int]bool, component *[]string) { //v - вершина
	visited[v] = true
	*component = append(*component, v)

	for _, neighbor := range graph[v] {
		if !visited[neighbor] {
			Dfs(graph, neighbor, visited, component)
		}
	}
}
