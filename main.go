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
		fmt.Println("error: check the arguments again")
		return
	}
	input, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("error in reading file")
		return
	}
	defer input.Close()
	scanner := bufio.NewScanner(input)
	Graph, err := funcs.ParseInput(scanner)
	if err != nil {
		fmt.Println("Error : Check your input again")
	}
	
}
