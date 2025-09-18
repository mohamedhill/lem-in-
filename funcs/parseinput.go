package funcs

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

type Coordinates struct {
	X int
	Y int
}

// a function to parse the input and verify format
func ParseInput(scanner *bufio.Scanner) (*Graph, string) {
	Graph := &Graph{Rooms: make(map[string]*Room)} // making a variable type pointer to Graph
	AntsRead := false
	var start, end bool
	fileinfo := ""
	coords := make(map[Coordinates]string)
	for scanner.Scan() {
		line := scanner.Text()
		fileinfo += line + "\n"
		line = strings.TrimSpace(line)
		// Ignoring Emptylines and comments

		if !AntsRead && line != "" && !strings.HasPrefix(line, "#") {
			Ants, err := strconv.Atoi(line)
			if Ants <= 0 || err != nil {
				fmt.Println("Error: Check the number of Ants; must be a positive integer number")
				return nil, ""

			}
			AntsRead = true
			Graph.Ants = Ants
			continue
		}
		if line == "" || strings.HasPrefix(line, "#") && !strings.HasPrefix(line, "##") {
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
				return nil, ""
			}
		}

		if strings.Contains(line, "-") {
			parts := strings.Split(line, "-")
			if len(parts) != 2 {
				fmt.Println("error in input check again; usage of links: example1-example2")
				return nil, ""
			}
			room1, exist1 := Graph.Rooms[parts[0]]
			room2, exist2 := Graph.Rooms[parts[1]]
			if !exist1 || !exist2 {
				fmt.Println("error room can't be linked (dosen't exist)")
				return nil, ""
			}
			for _, link := range room1.Links {
				if link == room2 {
					fmt.Println("ERROR: duplicate tunnel betwen ",parts[0],"and",parts[1])
					return nil, ""
				}
			}
			room1.Links = append(room1.Links, room2)
			room2.Links = append(room2.Links, room1)
		} else {
			splitted := strings.Fields(line)
			if len(splitted) != 3 {
				fmt.Println("error in input check again; usage of rooms: roomname x y")
				return nil, ""
			}
			if strings.HasPrefix(splitted[0], "#") || strings.HasPrefix(splitted[0], "L") {
				fmt.Println("error in input check again; room name can't start with L or #")
				return nil, ""
			}
			cordinate1, err := strconv.Atoi(splitted[1])
			if err != nil {
				fmt.Println("error in input checkroom cordinates")
				return nil, ""
			}
			cordinate2, err := strconv.Atoi(splitted[2])
			if err != nil {
				fmt.Println("error in input checkroom cordinates")
				return nil, ""
			}
			for _, v := range Graph.Rooms {
				if splitted[0] == v.Name {
					fmt.Println("Error Duplicate room names")
					return nil, ""
				}
			}
			coord := Coordinates{X: cordinate1, Y: cordinate2}
			if existingRoom, exists := coords[coord]; exists {
				fmt.Printf("Error: Duplicate coordinates for rooms '%s' and '%s'\n", existingRoom, splitted[0])
				return nil, ""
			}
			coords[coord] = splitted[0]

			room := &Room{Name: splitted[0], X: cordinate1, Y: cordinate2}
			if start {
				if end {
					fmt.Println("ERROR: the next roomt after the start flag undefind")
					return nil, ""
				}
				if Graph.Start != nil {
					fmt.Println("ERROR: multiple ##start definitions")
					return nil, ""
				}
				Graph.Start = room
				room.IsStart = true
				start = false
			}
			if end {
				if start {
					fmt.Println("ERROR: the next roomt after the start flag undefind")
					return nil, ""
				}
				if Graph.End != nil {
					fmt.Println("ERROR: multiple ##end definitions")
					return nil, ""
				}
				Graph.End = room
				room.IsEnd = true
				end = false
			}

			Graph.Rooms[splitted[0]] = room

		}
	}
	return Graph, fileinfo
}
