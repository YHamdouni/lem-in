package lemin

import (
	"fmt"
	"strings"
)

func Printfinal(DISTRIBUTION [][]string, data [][]string) {
	var final_result [][]string
	final_result = append(final_result, data...)
	final_result = append(final_result, DISTRIBUTION...)
	var count int
	for _, round := range final_result {
		count++
		if count == len(data) {
			fmt.Print("\n")
		}
		fmt.Println(strings.Join(round, " "))
	}
}
