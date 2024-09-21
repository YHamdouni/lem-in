package main

import (
	"fmt"
	"os"

	Functions "lemin/Functions"
)

func main() {
	Args := os.Args[1:]
	if len(Args) != 1 {
		fmt.Println("You didn't enter a file name, or you entered more than one file.")
		os.Exit(0)
	}
	file := Args[0]
	data, err := Functions.Readfile(file)
	if err != nil {
		fmt.Println("error reading file")
		os.Exit(0)
	}
	graph, start, end := Functions.RoomsDetails(data)
	Allpath := Functions.WAYS(graph, start, end)
	fmt.Println("All paths:", Allpath)
}
