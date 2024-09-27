package lemin

import "fmt"

func FinalResult(paths [][]string, totalAnts int, start string, end string) [][]string {
	type Ant struct {
		ID        int
		PathIndex int
		Position  int
	}

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
			if ants[i].PathIndex != -1 {
				if ants[i].Position+1 < len(paths[ants[i].PathIndex]) {
					ants[i].Position++
					currentRoom := paths[ants[i].PathIndex][ants[i].Position]
					if currentRoom != end {
						if !occupied[currentRoom] {
							moves = append(moves, fmt.Sprintf("L%d-%s", ants[i].ID, currentRoom))
							occupied[currentRoom] = true
						}
					} else {
						antsAtEnd++
					}
				} else {
					// Ant has reached the end
					antsAtEnd++
				}
			}
		}

		// Put new ants on paths
		for pathIndex, path := range paths {
			if currentAnt < totalAnts && ants[currentAnt].PathIndex == -1 && len(path) > 1 {
				nextRoom := path[1] // First room after start
				if !occupied[nextRoom] {
					ants[currentAnt].PathIndex = pathIndex
					ants[currentAnt].Position = 1
					moves = append(moves, fmt.Sprintf("L%d-%s", ants[currentAnt].ID, nextRoom))
					occupied[nextRoom] = true
					currentAnt++
				}
			}
		}

		if len(moves) > 0 {
			rounds = append(rounds, moves)
		}
	}

	return rounds
}

// //////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// /////// chat gpt + claude best result for now
// /////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// func FinalResult(paths [][]string, totalAnts int, start string, end string) [][]string {
// 	type Ant struct {
// 		ID        int
// 		PathIndex int
// 		Position  int
// 	}

// 	// Initialize ants
// 	ants := make([]Ant, totalAnts)
// 	for i := range ants {
// 		ants[i] = Ant{ID: i + 1, PathIndex: i % len(paths), Position: -1} // -1 means at start
// 	}

// 	var rounds [][]string
// 	occupied := make(map[string]bool)

// 	for {
// 		var moves []string
// 		antsAtEnd := 0
// 		movedThisRound := false

// 		// Process each ant
// 		for i := range ants {
// 			ant := &ants[i]
// 			// fmt.Println("ant", ant)
// 			if ant.Position == len(paths[ant.PathIndex])-1 {
// 				antsAtEnd++
// 				continue
// 			}

// 			nextPos := ant.Position + 1
// 			nextRoom := paths[ant.PathIndex][nextPos]

// 			if nextRoom == end || (!occupied[nextRoom]) {
// 				if ant.Position >= 0 {
// 					occupied[paths[ant.PathIndex][ant.Position]] = false
// 				}
// 				ant.Position = nextPos
// 				if nextRoom != end {
// 					occupied[nextRoom] = true
// 				}
// 				if nextRoom != start {
// 					moves = append(moves, fmt.Sprintf("L%d-%s", ant.ID, nextRoom))
// 					// fmt.Println("move", moves)
// 				}
// 				movedThisRound = true
// 			}
// 		}
// 		if movedThisRound {
// 			rounds = append(rounds, moves)
// 		} else {
// 			// If no moves were made, we need to break to avoid infinite loop
// 			break
// 		}

// 		if antsAtEnd == totalAnts {
// 			break
// 		}
// 	}

// 	return rounds[1:]
// }

/////////////////////////////////////////////////////////////////////////
//claude ai
////////////////////////////////////////////////////////////////////////

// func FinalResult(paths [][]string, totalAnts int, start string, end string) [][]string {
// 	type Ant struct {
// 		ID        int
// 		PathIndex int
// 		Position  int
// 	}

// 	// Initialize ants
// 	ants := make([]Ant, totalAnts)
// 	for i := range ants {
// 		ants[i] = Ant{ID: i + 1, PathIndex: (i) % len(paths), Position: -1} // -1 means at start
// 	}

// 	var rounds [][]string
// 	occupied := make(map[string]bool)

// 	for {
// 		var moves []string
// 		antsAtEnd := 0
// 		// Process each ant
// 		for i := range ants {
// 			ant := &ants[i]
// 			if ant.Position == len(paths[ant.PathIndex])-1 {
// 				antsAtEnd++
// 				continue
// 			}

// 			nextPos := ant.Position + 1
// 			nextRoom := paths[ant.PathIndex][nextPos]

// 			if !occupied[nextRoom] || nextRoom == end {
// 				if ant.Position >= 0 {
// 					occupied[paths[ant.PathIndex][ant.Position]] = false
// 				}
// 				ant.Position = nextPos
// 				if nextRoom != end {
// 					occupied[nextRoom] = true
// 				}
// 				moves = append(moves, fmt.Sprintf("L%d-%s", ant.ID, nextRoom))
// 			}
// 		}

// 		if len(moves) > 0 {
// 			rounds = append(rounds, moves)
// 			fmt.Println("round", rounds)
// 		}

// 		if antsAtEnd == totalAnts {
// 			break
// 		}
// 	}

// 	return rounds
// }
