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

	antsPerPath := distributeAntsOptimally(paths, totalAnts)
	ants := make([]Ant, totalAnts)
	for i := range ants {
		ants[i] = Ant{ID: i + 1, PathIndex: -1, Position: -1}
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
				} else {
					occupied[currentRoom] = true
				}
				moves = append(moves, fmt.Sprintf("L%d-%s", ants[i].ID, currentRoom))
			}
		}

		// Put new ants on paths
		for pathIndex, path := range paths {
			if antsPerPath[pathIndex] > 0 && currentAnt < totalAnts && !occupied[path[1]] {
				ants[currentAnt].PathIndex = pathIndex
				ants[currentAnt].Position = 1
				moves = append(moves, fmt.Sprintf("L%d-%s", ants[currentAnt].ID, path[1]))
				if path[1] != end {
					occupied[path[1]] = true
				}
				currentAnt++
				antsPerPath[pathIndex]--
			}
		}

		if len(moves) > 0 {
			rounds = append(rounds, moves)
		} else {
			break
		}
	}
	return rounds
}

// Optimized function to distribute ants across paths
func distributeAntsOptimally(paths [][]string, totalAnts int) []int {
	type PathInfo struct {
		index  int
		length int
		ants   int
	}

	pathInfos := make([]PathInfo, len(paths))
	for i, path := range paths {
		pathInfos[i] = PathInfo{index: i, length: len(path), ants: 0}
	}
	// fmt.Println("Pathinfos", pathInfos)

	sort.Slice(pathInfos, func(i, j int) bool {
		return pathInfos[i].length < pathInfos[j].length
	})

	for totalAnts > 0 {
		bestIndex := 0
		bestTime := float64(pathInfos[0].length + pathInfos[0].ants)
		for i := 1; i < len(pathInfos); i++ {
			time := float64(pathInfos[i].length + pathInfos[i].ants)
			if time < bestTime {
				bestTime = time
				bestIndex = i
			}
		}
		pathInfos[bestIndex].ants++
		totalAnts--
	}
	// fmt.Println("Pathinfos", pathInfos)

	antsPerPath := make([]int, len(paths))
	for _, pi := range pathInfos {
		antsPerPath[pi.index] = pi.ants
	}

	return antsPerPath
}
