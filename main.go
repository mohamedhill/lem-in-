package main

import "fmt"

func Dfs(Graph map[string][]string, node string, visited map[string]bool) {
	visited[node] = true
	fmt.Println(node)
	for _, v := range Graph[node] {
		if !visited[v] {
			visited[v] = true
			Dfs(Graph, v, visited)
		}
	}
}

func main() {
	Graph := map[string][]string{
		"A": {"B"},
		"B": {"C", "D"},
		"C": {"A"},
		"D": {"E"},
		"E": {"B"},
		"F": {"E"},
	}
	Bfs(Graph, "A")
	fmt.Println("--------------------------------")
	visited := make(map[string]bool)

	Dfs(Graph, "A", visited)
}

func Bfs(Graph map[string][]string, start string) {
	visited := make(map[string]bool)
	var queue []string
	queue = append(queue, start)
	visited[start] = true
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		fmt.Println(node)

		for _, v := range Graph[node] {
			if !visited[v] {
				visited[v] = true
				queue = append(queue, v)
			}
		}
	}
}
