package lemin

import "sort"

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
