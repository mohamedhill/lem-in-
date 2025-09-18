package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"lem-in/funcs"
)

func main() {
	if len(os.Args) != 2 || !strings.HasSuffix(os.Args[1], ".txt") {
		fmt.Println("error: Check your Args; Usage: go run . test.txt")
		return
	}

	input, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("error: can't open file", err)
		return
	}
	defer input.Close()
	scanner := bufio.NewScanner(input)
	Graph, fileinfo := funcs.ParseInput(scanner)

	if Graph == nil || Graph.Start == nil || Graph.End == nil {
		fmt.Println("ERROR: invalid data format")
		return
	}

	var bestPaths [][]string
	var allPaths [][]string
	if len(Graph.Rooms) > 30 {
		allPaths = funcs.FindPathsBFS(Graph)
		if allPaths == nil {
			allPaths = funcs.FindAllPaths(Graph)
			if allPaths == nil {
				fmt.Println("no valid path found")
				return
			}
		}
	} else {
		allPaths = funcs.FindAllPaths(Graph)
		if allPaths == nil {
			allPaths = funcs.FindPathsBFS(Graph)

			if allPaths == nil {
				fmt.Println("ERROR: no valid path found DFS")
				return
			}
		}
	}
	fmt.Println("allpaths here:", allPaths)
	bestPaths = funcs.FilterPaths(allPaths, Graph.Ants)
	if len(bestPaths) == 0 {
		fmt.Println("ERROR: no non-overlapping paths found")
		return
	}

	antDistribution := funcs.DistributeAnts(bestPaths, Graph.Ants)
	result, _ := funcs.SimulateAntMovement(bestPaths, antDistribution)
	if len(fileinfo) != 0 {
		fmt.Print(fileinfo)
	}
	fmt.Print(result)
}
