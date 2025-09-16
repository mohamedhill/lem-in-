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
	Graph, fileinfo:= funcs.ParseInput(scanner)

	if Graph == nil || Graph.Start == nil || Graph.End == nil {
		fmt.Println("ERROR: invalid data format")
		return
	}

	var bestPaths [][]string
	if Graph.Ants > 20 || len(Graph.Rooms) > 50 {
		bestPaths = funcs.FindPathsBFS(Graph)
		if bestPaths==nil{
			fmt.Println("ERROR: no valid path found")
			return
		} 
	} else {
		allPaths := funcs.FindAllPaths(Graph)
		if len(allPaths) == 0 {
			fmt.Println("ERROR: no valid path found")
			return
		}
		bestPaths = funcs.FilterPaths(allPaths)
		if len(bestPaths) == 0 {
			fmt.Println("ERROR: no non-overlapping paths found")
			return
		}
	}
	antDistribution := funcs.DistributeAnts(bestPaths, Graph.Ants)
	result, _ := funcs.SimulateAntMovement(bestPaths, antDistribution)
	if len(fileinfo)!=0{
		fmt.Print(fileinfo)
	}	
	fmt.Print(result)
}
