package lemin

import (
	"bufio"
	"fmt"
	"os"
)

func Readfile(filename string) ([][]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("error opening file, or file doesn't exist")
		os.Exit(0)
	}
	var result [][]string
	Scanner := bufio.NewScanner(file)
	for Scanner.Scan() {
		line := Scanner.Text()
		var slice []string
		slice = append(slice, line)
		result = append(result, slice)
	}
	if len(result) == 0 {
		fmt.Println("ERROR:\tthe file is empty")
		os.Exit(0)
	}
	return result, nil
}
