package lemin

import (
	"fmt"
	"strings"
)

func Printfinal(distribution [][]string, data [][]string) {
	// Print the input data
	for _, line := range data {
		fmt.Println(strings.Join(line, " "))
	}

	fmt.Print("\n") // Empty line between input and output

	// Print the ant movements
	for _, round := range distribution {
		fmt.Println(strings.Join(round, " "))
	}
}
