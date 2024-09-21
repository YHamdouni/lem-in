package lemin

import "strings"

func GenerateGraph(rooms [][]string) map[string][]string {
	graph1 := make(map[string][]string)
	graph2 := make(map[string][]string)
	for _, edges := range rooms {
		for _, edge := range edges {
			nodes := strings.Split(edge, "-")
			if len(nodes) != 2 {
				continue
			}
			node1 := nodes[0]
			node2 := nodes[1]
			graph1[node1] = append(graph1[node1], node2)
			graph2[node2] = append(graph2[node2], node1)
		}
	}
	combinedGraph := ConcatMaps(graph1, graph2)
	return combinedGraph
}
