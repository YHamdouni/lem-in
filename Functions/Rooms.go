package lemin

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Rooms(data [][]string) (map[int][]int, int, int) {
	var rooms [][]string
	check := false
	var count1 int
	var count2 int
	var start string
	var end string
	for _, lines := range data {
		for _, line := range lines {
			if count2 == 1 {
				end = string(line[0])
				count2 = 0
			}
			if count1 == 1 {
				start = string(line[0])
				count1 = 0
			}
			if len(line) == 3 {
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

	starting, err := strconv.Atoi(start)
	if err != nil {
		fmt.Println("there is error in converting the start")
		os.Exit(0)
	}
	Ending, err := strconv.Atoi(end)
	if err != nil {
		fmt.Println("there is error in converting the end")
		os.Exit(0)
	}
	rooms = Checkint(rooms)
	graph := Graph(rooms)
	return graph, starting, Ending
}

func Graph(rooms [][]string) map[int][]int {
	graph1 := make(map[int][]int)
	graph2 := make(map[int][]int)
	for _, edges := range rooms {
		for _, edge := range edges {
			nodes := strings.Split(edge, "-")
			if len(nodes) != 2 {
				continue
			}
			node1, err1 := strconv.Atoi(nodes[0])
			node2, err2 := strconv.Atoi(nodes[1])
			if err1 != nil || err2 != nil {
				continue
			}
			graph1[node1] = append(graph1[node1], node2)
			graph2[node2] = append(graph2[node2], node1)
		}
	}
	combinedGraph := concatMaps(graph1, graph2)
	return combinedGraph
}

func concatMaps(m1, m2 map[int][]int) map[int][]int {
	result := make(map[int][]int)
	for k, v := range m1 {
		result[k] = v
	}
	for k, v := range m2 {
		if existing, ok := result[k]; ok {
			result[k] = append(existing, v...)
		} else {
			result[k] = v
		}
	}
	return result
}

func Checkint(rooms [][]string) [][]string {
	var result [][]string
	check := false
	for _, s := range rooms {
		for _, v := range s {
			for _, g := range v {
				if g >= '0' && g <= '9' || g == ' ' {
					check = true
				} else {
					check = false
				}
			}
		}
		if check == true {
			result = append(result, s)
		}
	}
	return rooms
}
