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
	// fmt.Println("ggraph", graph)
	// fmt.Println("/////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////")
	Allpaths := Functions.WAYS(graph, start, end)
	ShortestPaths := Functions.GetBestPaths(Allpaths, numberOfAnts)
	fmt.Println("/////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////")
	fmt.Println("sho", ShortestPaths)
	sortedPaths := Functions.SortPaths(ShortestPaths)
	DISTRIBUTION := Functions.FinalResult(sortedPaths, numberOfAnts, start, end)
	Functions.Printfinal(DISTRIBUTION, data)
}
