package lemin

import (
	"fmt"
	"os"
	"strings"
)

func CheckArgs(Args []string) string {
	if len(Args) != 1 {
		fmt.Println("You didn't enter a file name, or you entered more than one file.")
		os.Exit(0)
	}
	file := Args[0]
	if !strings.HasSuffix(file, ".txt") {
		fmt.Println("use a file extention : = '.txt'")
		os.Exit(0)
	}
	return file
}
