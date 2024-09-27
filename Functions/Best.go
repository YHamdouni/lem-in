package lemin

import "sort"

func GetBestPaths(allPaths [][]string, numberOfAnts int) [][]string {
	type PathInfo struct {
		path   []string
		length int
	}

	// Convert paths to PathInfo and sort by length
	pathInfos := make([]PathInfo, len(allPaths))
	for i, path := range allPaths {
		pathInfos[i] = PathInfo{path: path, length: len(path) - 1} // -1 because we count edges, not nodes
	}
	sort.Slice(pathInfos, func(i, j int) bool {
		return pathInfos[i].length < pathInfos[j].length
	})

	// Helper function to check if two paths share any room (except start and end)
	sharesRoom := func(path1, path2 []string) bool {
		set := make(map[string]bool)
		for i := 1; i < len(path1)-1; i++ {
			set[path1[i]] = true
		}
		for i := 1; i < len(path2)-1; i++ {
			if set[path2[i]] {
				return true
			}
		}
		return false
	}

	// Helper function to calculate the number of rounds for a given set of paths
	calculateRounds := func(paths []PathInfo, ants int) int {
		if len(paths) == 0 {
			return 0
		}
		totalLength := 0
		for _, p := range paths {
			totalLength += p.length
		}
		return (totalLength + ants - 1) / len(paths)
	}

	bestPaths := []PathInfo{pathInfos[0]}
	bestRounds := calculateRounds(bestPaths, numberOfAnts)

	for i := 1; i < len(pathInfos); i++ {
		candidatePaths := make([]PathInfo, len(bestPaths))
		copy(candidatePaths, bestPaths)

		canAdd := true
		for _, p := range candidatePaths {
			if sharesRoom(p.path, pathInfos[i].path) {
				canAdd = false
				break
			}
		}

		if canAdd {
			candidatePaths = append(candidatePaths, pathInfos[i])
			candidateRounds := calculateRounds(candidatePaths, numberOfAnts)

			if candidateRounds < bestRounds {
				bestPaths = candidatePaths
				bestRounds = candidateRounds
			}
		}
	}

	result := make([][]string, len(bestPaths))
	for i, p := range bestPaths {
		result[i] = p.path
	}

	return result
}
