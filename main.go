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
	// fmt.Println("alll", Allpaths)
	// sortedPaths := Functions.SortPaths(Allpaths)
	ShortestPaths := Functions.GetBestPaths(Allpaths, numberOfAnts)
	fmt.Println("Paths used", ShortestPaths)
	DISTRIBUTION := Functions.FinalResult(ShortestPaths, numberOfAnts, start, end)
	Functions.Printfinal(DISTRIBUTION, data)
}
