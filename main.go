package main

import (
	"fmt"
	"os"

	Functions "lemin/Functions"
)

func main() {
	Args := os.Args[1:]
	file := Functions.CheckArgs(Args)
	data, err := Functions.Readfile(file)
	if err != nil {
		fmt.Println("error reading file")
		os.Exit(0)
	}
	graph, start, end, numberOfAnts := Functions.RoomsDetails(data)
	Allpaths := Functions.WAYS(graph, start, end)
	fmt.Println("len", len(Allpaths))
	// i need to remove this "numberOfAnts <= 0" when i will finish the project
	if numberOfAnts <= 0 {
		fmt.Println("ERROR: invalid data format")
		os.Exit(0)
	}
	////////////////////////////////////////////////////
	ShortestPaths := Functions.GetShortestPaths(Allpaths)
	sortedPaths := Functions.SortPaths(ShortestPaths)
	fmt.Println("Paths", sortedPaths)
	Functions.Print(sortedPaths, numberOfAnts, start, end)
}
