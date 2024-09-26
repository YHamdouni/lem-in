package lemin

import (
	"fmt"
	"strconv"
	"strings"
)

func Print(paths [][]string, totalAnts int, start string, end string) {
	positions := make([]int, totalAnts)
	occupied := make(map[string]int)
	// rounds := make([][]string, 0)
	// var result string
	for {
		for i := 0; i < totalAnts; i++ {
			pathIndex := i % len(paths)
			currentroom := paths[pathIndex][positions[i]]
			fmt.Println("currentroom", currentroom)
			positions[i]++
			occupied[currentroom]++
		}
	}
}

func FinalResult(paths [][]string, totalAnts int, start string, end string) {
	// Track the positions of each ant
	positions := make([]int, totalAnts)
	occupied := make(map[string]int) // Tracks which rooms are occupied
	rounds := make([][]string, 0)    // To store output for each round

	// Continue until all ants reach the end
	for {
		output := make([]string, 0)
		allDone := true // Flag to check if all ants have reached their destination

		// Loop through each ant
		for ant := 0; ant < totalAnts; ant++ {
			pathIndex := ant % len(paths)
			if positions[ant] < len(paths[pathIndex]) {
				// fmt.Println(positions[ant], len(paths[pathIndex]))
				currentRoom := paths[pathIndex][positions[ant]]

				// Check if the room is occupied (except start and end rooms)
				if currentRoom != start && currentRoom != end && occupied[currentRoom] > 0 {
					continue // Skip this ant's movement if the room is occupied
				}

				allDone = false // At least one ant is still moving

				// If not in the starting position, format the output
				if positions[ant] > 0 {
					output = append(output, "L"+strconv.Itoa(ant+1)+"-"+currentRoom)
				}

				// Mark the current room as occupied
				occupied[currentRoom]++

				// Move the ant to the next position
				positions[ant]++
			}
		}

		// If all ants are done, break the loop
		if allDone {
			break
		}

		// Store the current round's output
		if len(output) > 0 {
			rounds = append(rounds, output)
		}

		// Reset the occupancy for the next round
		for _, room := range output {
			roomID := strings.Split(room, "-")[1]
			occupied[roomID]--
		}
	}

	// Print the output in the desired format
	for _, round := range rounds {
		fmt.Println(strings.Join(round, " "))
	}
}
