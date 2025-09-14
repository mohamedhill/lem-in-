package funcs

import (
	"fmt"
	"strings"
)

// SimulateAntMovement prints the moves for all ants using multiple paths
func SimulateAntMovement(paths [][]string, antDistribution [][]int) (string, int) {
	var finalResult string
	type AntPosition struct {
		ant  int
		path int
		step int
	}

	var antPositions []AntPosition
	for pathIndex, ants := range antDistribution {
		for _, ant := range ants {
			antPositions = append(antPositions, AntPosition{ant, pathIndex, 0})
		}
	}

	moveCount := 0
	for len(antPositions) > 0 {
		var moves []string
		var newPositions []AntPosition
		occupied := make(map[string]bool)
		startMoved := make(map[int]bool)

		for _, pos := range antPositions {
			if pos.step < len(paths[pos.path])-1 {
				nextRoom := paths[pos.path][pos.step+1]
				if pos.step == 0 && startMoved[pos.path] {
					newPositions = append(newPositions, pos)
					continue
				}
				if nextRoom == paths[pos.path][len(paths[pos.path])-1] || !occupied[nextRoom] {
					moves = append(moves, fmt.Sprintf("L%d-%s", pos.ant, nextRoom))
					newPositions = append(newPositions, AntPosition{pos.ant, pos.path, pos.step + 1})
					if nextRoom != paths[pos.path][0] && nextRoom != paths[pos.path][len(paths[pos.path])-1] {
						occupied[nextRoom] = true
					}
					if pos.step == 0 {
						startMoved[pos.path] = true
					}
				} else {
					newPositions = append(newPositions, pos)
				}
			}
		}
		if len(moves) > 0 {
			finalResult += strings.Join(moves, " ") + "\n"
		}
		antPositions = newPositions
		moveCount++
	}
	return finalResult, moveCount - 1
}
