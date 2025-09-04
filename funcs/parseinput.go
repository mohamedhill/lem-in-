package funcs

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

func ParseInput(scanner *bufio.Scanner) (*Graph, error) {
	Graph := &Graph{Rooms: make(map[string]*Room)}
	Line := -1
	var start, end bool
	for scanner.Scan() {
		Line++
		line := scanner.Text()
		line = strings.TrimSpace(line)
		// Ignoring Emptylines and comments 
		if line == "" || !strings.HasPrefix(line, "##") && strings.HasPrefix(line, "#") {
			continue
		}
		if strings.HasPrefix(line, "##") {
			if strings.TrimPrefix(line, "##") == "start" {
				start = true
			} else if strings.TrimPrefix(line, "##") == "end" {
				end = true
			} else {
				fmt.Println("error in input check again; usage: ##start||##end")
				return nil, nil
			}
			if Line == 0 {
				Ants, err := strconv.Atoi(line)
				if Ants <= 0 || err != nil {
					fmt.Println("Error: Check the number of Ants; must be a positive integer number")
					return nil, nil
				}
				Graph.Ants = Ants
			}
			if Line > 0 {
				if strings.Contains(line, "-") {
					parts := strings.Split(line, "-")
					if len(parts) != 2 {
						fmt.Println("error in input check again; usage of links: example1-example2")
						return nil, nil
					}

				} else {
					splitted := strings.Fields(line)
					if len(splitted) != 3 {
						fmt.Println("error in input check again; usage of rooms: roomname x y")
						return nil, nil
					}
					if strings.HasPrefix(splitted[0], "#") || strings.HasPrefix(splitted[0], "L") {
						fmt.Println("error in input check again; room name can't start with L or #")
						return nil, nil
					}
					x, err := strconv.Atoi(splitted[1])
					if err != nil {
						fmt.Println("error in input checkroom cordinates")
					}
					y, err := strconv.Atoi(splitted[2])
					if err != nil {
						fmt.Println("error in input checkroom cordinates")
					}
					if

				}
			}
		}
	}
	return Graph, nil
}
