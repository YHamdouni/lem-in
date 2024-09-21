package lemin

import (
	"strings"
)

func RoomsDetails(data [][]string) (map[string][]string, string, string) {
	var rooms [][]string
	check := false
	var count1 int
	var count2 int
	var start []string
	var end []string
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
				count2++
			} else if line == "##start" {
				count1++
			}
		}
	}
	graph := GenerateGraph(rooms)
	return graph, start[0], end[0]
}

