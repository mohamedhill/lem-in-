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

		// Ignoring Emptylines and comments
		if strings.TrimSpace(line) == "" || strings.HasPrefix(line, "#") && !strings.HasPrefix(line, "##") {
			continue
		}
		if strings.HasPrefix(line, "##") {
			if strings.TrimPrefix(line, "##") == "start" {
				start = true
				continue
			} else if strings.TrimPrefix(line, "##") == "end" {
				end = true
				continue
			} else {
				fmt.Println("error in input check again; usage: ##start||##end")
				return nil, nil
			}
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
				room1, exist1 := Graph.Rooms[parts[0]]
				room2, exist2 := Graph.Rooms[parts[1]]
				if !exist1 || !exist2 {
					fmt.Println("error room can't be linked (dosen't exist)")
					return nil, nil
				}
				room1.Links = append(room1.Links, room2)
				room2.Links = append(room2.Links, room1)
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
				cordinate1, err := strconv.Atoi(splitted[1])
				if err != nil {
					fmt.Println("error in input checkroom cordinates")
				}
				cordinate2, err := strconv.Atoi(splitted[2])
				if err != nil {
					fmt.Println("error in input checkroom cordinates")
					return nil, nil
				}
				for _, v := range Graph.Rooms {
					if splitted[0] == v.Name {
						fmt.Println("Error Duplicate room names")
						return nil, nil
					}
				}
				room := &Room{Name: splitted[0], X: cordinate1, Y: cordinate2}
				if start {
					Graph.Start = room
					start = false
					room.IsStart = true
				}
				if end {
					Graph.End = room
					end = false
					room.IsEnd = true
				}
				Graph.Rooms[splitted[0]] = room

			}
		}
	}
	return Graph, nil
}
