package lemin

type PathInfo struct {
	path   []string
	length int
}

// GetBestPaths finds the optimal paths for a given number of ants.
func GetBestPaths(allPaths [][]string, numberOfAnts int) [][]string {
	// Convert the provided paths into PathInfo structs, calculating their lengths
	pathInfos := make([]PathInfo, len(allPaths))
	for i, path := range allPaths {
		pathInfos[i] = PathInfo{path: path, length: len(path) - 1} // -1 because we count edges, not nodes
	}
	// sortPathsByLength(pathInfos)
	// Start with the shortest path as the best candidate
	bestPaths := []PathInfo{pathInfos[0]}
	// Helper function to calculate the number of rounds needed for a given set of paths
	// hadi katjib l path and katjib ch7al fih man move bach yowssal l end
	bestRounds := calculateRounds(bestPaths, numberOfAnts)

	// Iterate through the remaining paths to find the best combination
	for i := 1; i < len(pathInfos); i++ {
		candidatePaths := make([]PathInfo, len(bestPaths))
		copy(candidatePaths, bestPaths)
		canAdd := true
		// Check if the current path shares any room with the best paths
		for _, p := range candidatePaths {
			if sharesRoom(p.path, pathInfos[i].path) {
				canAdd = false // Cannot add this path due to overlap
				break
			}
		}
		// If it doesn't overlap, calculate the candidate rounds
		if canAdd {
			candidatePaths = append(candidatePaths, pathInfos[i])
			candidateRounds := calculateRounds(candidatePaths, numberOfAnts)

			// Update bestPaths if the new candidate has fewer rounds
			if candidateRounds < bestRounds {
				bestPaths = candidatePaths
				bestRounds = candidateRounds
			}
		}
	}
	// Prepare the result by extracting the paths from bestPaths
	result := make([][]string, len(bestPaths))
	for i, p := range bestPaths {
		result[i] = p.path
	}
	return result
}

// Helper function to check if two paths share any intermediate room (excluding start and end)
func sharesRoom(path1, path2 []string) bool {
	set := make(map[string]bool)
	for i := 1; i < len(path1)-1; i++ {
		set[path1[i]] = true
	}
	// Check if any intermediate room of path2 is in the set
	for i := 1; i < len(path2)-1; i++ {
		if set[path2[i]] {
			return true // A shared room is found
		}
	}
	return false // No shared rooms
}

func calculateRounds(pathlenghts []PathInfo, ants int) int {
	if len(pathlenghts) == 0 {
		return 0 // No paths means no rounds
	}
	totalLength := 0
	// Sum up the lengths of all paths
	for _, p := range pathlenghts {
		totalLength += p.length
	}
	// Calculate rounds based on total length and number of paths
	return (totalLength + ants - 1) / len(pathlenghts)
}
