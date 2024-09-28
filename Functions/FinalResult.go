package lemin

import (
	"fmt"
	"sort"
)

func FinalResult(paths [][]string, totalAnts int, start string, end string) [][]string {
	type Ant struct {
		ID        int
		PathIndex int
		Position  int
	}

	// Calculate the number of ants to send on each path
	antsPerPath := distributeAnts(paths, totalAnts)
	// antsPerPath := []int{1000, 20}

	// Initialize ants
	ants := make([]Ant, totalAnts)
	for i := range ants {
		ants[i] = Ant{ID: i + 1, PathIndex: -1, Position: -1} // -1 means not on a path yet
	}

	var rounds [][]string
	antsAtEnd := 0
	currentAnt := 0

	for antsAtEnd < totalAnts {
		var moves []string
		occupied := make(map[string]bool)

		// Move ants already on paths
		for i := range ants {
			if ants[i].PathIndex != -1 && ants[i].Position < len(paths[ants[i].PathIndex])-1 {
				ants[i].Position++
				currentRoom := paths[ants[i].PathIndex][ants[i].Position]
				if currentRoom == end {
					antsAtEnd++
					moves = append(moves, fmt.Sprintf("L%d-%s", ants[i].ID, currentRoom))
					occupied[currentRoom] = true
				} else {
					if !occupied[currentRoom] {
						moves = append(moves, fmt.Sprintf("L%d-%s", ants[i].ID, currentRoom))
						occupied[currentRoom] = true
					}
				}
			}
		}

		// Put new ants on paths
		for pathIndex, path := range paths {
			for antsPerPath[pathIndex] > 0 && currentAnt < totalAnts && !occupied[path[1]] {
				ants[currentAnt].PathIndex = pathIndex
				ants[currentAnt].Position = 1
				moves = append(moves, fmt.Sprintf("L%d-%s", ants[currentAnt].ID, path[1]))
				occupied[path[1]] = true
				currentAnt++
				antsPerPath[pathIndex]--
			}
		}
		if len(moves) > 0 {
			rounds = append(rounds, moves)
		} else if antsAtEnd < totalAnts {
			break
		}
	}
	return rounds
}

// Helper function to distribute ants across paths
func distributeAnts(paths [][]string, totalAnts int) []int {
	antsPerPath := make([]int, len(paths))
	remainingAnts := totalAnts

	// Sort paths by length (shortest first)
	sort.Slice(paths, func(i, j int) bool {
		return len(paths[i]) < len(paths[j])
	})

	// Distribute ants favoring shorter paths
	for remainingAnts > 0 {
		for i := range paths {
			if remainingAnts > 0 {
				antsPerPath[i]++
				remainingAnts--
			} else {
				break
			}
		}
	}
	
	return antsPerPath
}
