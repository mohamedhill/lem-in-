package funcs

func PathFinding(graph *Graph) []string {
	queue := [][]*Room{{graph.Start}}
	visited := make(map[*Room]bool)

	for len(queue) > 0 {
		path := queue[0]
		queue = queue[1:]

		currntPath := path[len(path)-1]
		if currntPath == graph.End {
			var result []string
			for _, r := range path {
				result = append(result, r.Name)
			}
			return result
		}
		if visited[currntPath] {
			continue
		}
		visited[currntPath] = true
		for _, neighbor := range currntPath.Links {
			if !visited[neighbor] {
				newPath := append([]*Room{}, path...)
				newPath = append(newPath, neighbor)
				queue = append(queue, newPath)
			}
		}
	}
	return nil
}
