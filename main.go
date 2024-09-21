package main

import (
	"fmt"
	"os"

	Functions "lemin/Functions"
)

func WAYS(graph map[int][]int, start int, end int) [][]int {
	var result [][]int
	var path []int
	result = dfs(graph, start, end, path)
	return result
}

func dfs(graph map[int][]int, current int, end int, path []int) [][]int {
	var Paths [][]int
	path = append(path, current)
	if current == end {
		tempPath := make([]int, len(path))
		CopySlice(tempPath, path)
		Paths = append(Paths, tempPath)
		return Paths
	}
	for _, neighbor := range graph[current] {
		if !contains(path, neighbor) {
			Paths = append(Paths, dfs(graph, neighbor, end, path)...)
		}
	}

	return Paths
}

func CopySlice(dest, src []int) {
	n := len(dest)
	for i := 0; i < n; i++ {
		dest[i] = src[i]
	}
}

func contains(path []int, node int) bool {
	for _, n := range path {
		if n == node {
			return true
		}
	}
	return false
}

func main() {
	data, err := Functions.Readfile("file.txt")
	if err != nil {
		fmt.Println("error reading file")
		os.Exit(0)
	}
	graph, start, end := Functions.Rooms(data)
	PathOf := WAYS(graph, start, end)
	fmt.Println("start to end", PathOf)
}
