package lemin

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func RoomsDetails(data [][]string) (map[string][]string, string, string, int) {
	var rooms [][]string
	check := false
	checkend := true
	checkstart := true
	var count1 int
	var count2 int
	var start []string
	var end []string
	numberOfAnts, err := strconv.Atoi(data[0][0])
	if err != nil || numberOfAnts <= 0 {
		fmt.Println("ERROR: invalid data format, invalid number of Ants")
		os.Exit(0)
	}
	for _, lines := range data {
		for _, line := range lines {
			split := strings.Split(line, "-")
			if count2 == 1 {
				end = strings.Split(line, " ")
				count2 = 0
			}
			if count1 == 1 {
				start = strings.Split(line, " ")
				count1 = 0
			}
			if len(split) == 2 {
				check = true
			}
			if check {
				rooms = append(rooms, lines)
			}
			check = false
			if line == "##end" {
				checkend = false
				count2++
			} else if line == "##start" {
				count1++
				checkstart = false
			}
		}
	}

	if checkend {
		fmt.Println("ERROR: invalid data format, no end room found")
		os.Exit(0)
	} else if checkstart {
		fmt.Println("ERROR: invalid data format, no start room found")
		os.Exit(0)
	} else if start[0] == "L" || start[0] == "#" {
		fmt.Println("ERROR: invalid data format, no start room found")
		os.Exit(0)
	}

	graph := GenerateGraph(rooms)
	return graph, start[0], end[0], numberOfAnts
}
