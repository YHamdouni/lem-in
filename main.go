package main

import (
	"fmt"
	"os"
	"strings"

	Functions "lemin/Functions"
)

func main() {
	Args := os.Args[1:]
	if len(Args) != 1 {
		fmt.Println("You didn't enter a file name, or you entered more than one file.")
		os.Exit(0)
	}
	file := Args[0]
	if !strings.HasSuffix(file, ".txt") {
		fmt.Println("use a file extention : = '.txt'")
		os.Exit(0)
	}
	data, err := Functions.Readfile(file)
	if err != nil {
		fmt.Println("error reading file")
		os.Exit(0)
	}
	graph, start, end, numberOfAnts := Functions.RoomsDetails(data)
	Allpaths := Functions.WAYS(graph, start, end)
	// i need to remove this "numberOfAnts <= 0 " when i will finish the project
	if numberOfAnts <= 0 || len(Allpaths) == 0 {
		fmt.Println("ERROR: invalid data format")
		os.Exit(0)
	}
	fmt.Println("graph", graph)
	ShortestWhitoutSort := Functions.GetShortestPaths(Allpaths)
	fmt.Println("no sorted", ShortestWhitoutSort)
	Functions.SortPaths(ShortestWhitoutSort)
	ShortestPaths := Functions.GetShortestPaths(ShortestWhitoutSort)
	fmt.Println("yessorted", ShortestPaths)
	fmt.Println("///////////////////////")
	if len(ShortestWhitoutSort) > len(ShortestPaths) {
		fmt.Println("no sort", ShortestWhitoutSort)
	} else {
		fmt.Println("sorted", ShortestPaths)
	}
}

// bestPaths := Functions.SortPaths(Allpath)
// for _, path := range bestPaths {
// 	fmt.Println(path)
// }
