package lemin

import (
	"math"
	"sort"
)

type PathInfo struct {
	path   []string
	length int
}

func GetBestPaths(allPaths [][]string, numberOfAnts int) [][]string {
	if len(allPaths) == 0 || numberOfAnts <= 0 {
		return nil
	}

	pathInfos := make([]PathInfo, len(allPaths))
	for i, path := range allPaths {
		pathInfos[i] = PathInfo{path: path, length: len(path) - 1}
	}

	sort.Slice(pathInfos, func(i, j int) bool {
		return pathInfos[i].length < pathInfos[j].length
	})

	bestPaths := []PathInfo{}
	bestRounds := math.MaxInt32

	for i := 0; i < len(pathInfos); i++ {
		candidatePaths := []PathInfo{pathInfos[i]}
		for j := i + 1; j < len(pathInfos); j++ {
			if !hasOverlap(candidatePaths, pathInfos[j]) {
				candidatePaths = append(candidatePaths, pathInfos[j])
			}
		}

		rounds := calculateRoundsEfficient(candidatePaths, numberOfAnts)
		if rounds < bestRounds {
			bestRounds = rounds
			bestPaths = make([]PathInfo, len(candidatePaths))
			copy(bestPaths, candidatePaths)
		}

		// Early stopping condition
		if bestRounds <= pathInfos[i].length {
			break
		}
	}

	result := make([][]string, len(bestPaths))
	for i, p := range bestPaths {
		result[i] = p.path
	}
	return result
}

func hasOverlap(paths []PathInfo, newPath PathInfo) bool {
	set := make(map[string]bool)
	for _, path := range paths {
		for i := 1; i < len(path.path)-1; i++ {
			set[path.path[i]] = true
		}
	}
	for i := 1; i < len(newPath.path)-1; i++ {
		if set[newPath.path[i]] {
			return true
		}
	}
	return false
}

func calculateRoundsEfficient(paths []PathInfo, ants int) int {
	if len(paths) == 0 {
		return 0
	}

	// Create a copy of paths to avoid modifying the original
	pathsCopy := make([]PathInfo, len(paths))
	for i, p := range paths {
		pathsCopy[i] = PathInfo{path: p.path, length: p.length}
	}

	// Distribute ants optimally among paths
	for ants > 0 {
		shortestIndex := 0
		for i := 1; i < len(pathsCopy); i++ {
			if pathsCopy[i].length < pathsCopy[shortestIndex].length {
				shortestIndex = i
			}
		}
		pathsCopy[shortestIndex].length++
		ants--
	}

	maxLength := 0
	for _, path := range pathsCopy {
		if path.length > maxLength {
			maxLength = path.length
		}
	}

	return maxLength - 1
}
