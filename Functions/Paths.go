package lemin

func WAYS(graph map[string][]string, start string, end string) [][]string {
	var result [][]string
	var path []string
	result = dfs(graph, start, end, path)
	return result
}

func dfs(graph map[string][]string, current string, end string, path []string) [][]string {
	var Paths [][]string
	path = append(path, current)
	if current == end {
		tempPath := make([]string, len(path))
		CopySlice(tempPath, path)
		Paths = append(Paths, tempPath)
		return Paths
	}
	for _, neighbor := range graph[current] {
		if !Contains(path, neighbor) {
			// Call dfs recursively on the neighbor to explore further paths
			foundPaths := dfs(graph, neighbor, end, path)

			// Add all found paths from this call to the main Paths slice
			Paths = append(Paths, foundPaths...)
		}
	}

	return Paths
}
