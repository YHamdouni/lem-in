package lemin

import (
	"sort"
)

func CopySlice(dest, src []string) {
	n := len(dest)
	for i := 0; i < n; i++ {
		dest[i] = src[i]
	}
}

func Contains(path []string, node string) bool {
	for _, n := range path {
		if n == node {
			return true
		}
	}
	return false
}

func SortPaths(paths [][]string) [][]string {
	sort.Slice(paths, func(i, j int) bool {
		return len(paths[i]) < len(paths[j])
	})
	return paths
}

// func sortPathsByLengthstring(paths []string) []string {
// 	n := len(paths)
// 	for i := 0; i < n-1; i++ {
// 		for j := 0; j < n-i-1; j++ {
// 			if len(paths[j]) > len(paths[j+1]) {
// 				// Échange
// 				paths[j], paths[j+1] = paths[j+1], paths[j]
// 			}
// 		}
// 	}
// 	return paths
// }

func sortPathsByLength(paths []PathInfo) []PathInfo {
	n := len(paths)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if paths[j].length > paths[j+1].length {
				// Swap
				paths[j], paths[j+1] = paths[j+1], paths[j]
			}
		}
	}
	return paths
}
