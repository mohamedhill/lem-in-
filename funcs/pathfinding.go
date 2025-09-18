package funcs

import (
	
	"sort"
)

// FindAllPaths returns all paths from start to end as [][]string
func FindAllPaths(graph *Graph) [][]string {
	var allPaths [][]string
	var dfs func(current *Room, path []string, visited map[*Room]bool)
	dfs = func(current *Room, path []string, visited map[*Room]bool) {
		if current == graph.End {
			result := append([]string{}, path...)
			allPaths = append(allPaths, result)
			return
		}
		for _, neighbor := range current.Links {
			if !visited[neighbor] {
				visited[neighbor] = true
				dfs(neighbor, append(path, neighbor.Name), visited)
				visited[neighbor] = false
			}
		}
	}
	visited := map[*Room]bool{graph.Start: true}
	dfs(graph.Start, []string{graph.Start.Name}, visited)

	return allPaths
}

// a function to filter paths by removing overlapping ones
func FilterPaths(allPaths [][]string,antnum int) [][]string  {
	var oneantpath [][]string
	sort.Slice(allPaths, func(i, j int) bool {
		return len(allPaths[i]) < len(allPaths[j])
	})
	if antnum ==1{
		oneantpath=append(oneantpath, allPaths[0])
		return oneantpath
	}
	
	var bestPath [][]string
	var check func(current [][]string, left [][]string)
	check = func(current [][]string, left [][]string) {
		if len(current) > len(bestPath) {
			bestPath = make([][]string, len(current))
			copy(bestPath, current)
		}
		for i, path := range left {
			if canAddPath(current, path) {
				newCurrent := append([][]string{}, current...)
				newCurrent = append(newCurrent, path)
				newRemaining := [][]string{}
				for j, p := range left {
					if i != j && canAddPath(newCurrent, p) {
						newRemaining = append(newRemaining, p)
					}
				}
				check(newCurrent, newRemaining)
			}
		}
	}
	check([][]string{}, allPaths)
	return bestPath
}

// a functiion to avoid overlapping
func canAddPath(oldPath [][]string, newPath []string) bool {
	used := map[string]bool{}
	for _, path := range oldPath {
		for i := 1; i < len(path)-1; i++ {
			used[path[i]] = true
		}
	}
	for i := 1; i < len(newPath)-1; i++ {
		if used[newPath[i]] {
			return false
		}
	}
	return true
}

// a function to find shortest way using bfs
func FindPathsBFS(graph *Graph) [][]string {
	maxPaths := 20
	var paths [][]string
	queue := [][]*Room{{graph.Start}}

	for len(queue) > 0 && len(paths) < maxPaths {
		path := queue[0]
		queue = queue[1:]
		last := path[len(path)-1]

		if last == graph.End {
			var strPath []string
			for _, r := range path {
				strPath = append(strPath, r.Name)
			}
			paths = append(paths, strPath)
			continue
		}

		
		visitedInPath := map[*Room]bool{}
		for _, r := range path {
			visitedInPath[r] = true
		}

		for _, neighbor := range last.Links {
			if visitedInPath[neighbor] {
				continue 
			}
			newPath := append([]*Room{}, path...)
			newPath = append(newPath, neighbor)
			queue = append(queue, newPath)
		}
	}

	return paths
}

func DistributeAnts(paths [][]string, antCount int) [][]int {
	antDist := make([][]int, len(paths))
	pathLen := make([]int, len(paths))
	for i, p := range paths {
		pathLen[i] = len(p)
	}
	minTurns := 0
	// Number of ants that can arrive by turn minturn on a path of length l
	for {
		minTurns++
		sum := 0
		for _, l := range pathLen {
			if minTurns >= l {
				sum += minTurns - l + 1
			}
		}
		if sum >= antCount {
			break
		}

	}
	remaining := antCount
	antNum := 1
	for i, l := range pathLen {
		n := minTurns - l + 1
		if n > 0 {
			if n > remaining {
				n = remaining
			}
			for a := 0; a < n; a++ {
				antDist[i] = append(antDist[i], antNum)
				antNum++
				remaining--
			}
		}
	}
	for remaining > 0 {
		shortest := 0
		for i := 1; i < len(pathLen); i++ {
			if pathLen[i] < pathLen[shortest] {
				shortest = i
			}
		}
		antDist[shortest] = append(antDist[shortest], antNum)
		antNum++
		remaining--
	}
	return antDist
}
