package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) <= 1 || os.Args[1] == "" {
		fmt.Print(0)
		return
	}
	line := os.Args[1]
	words := strings.Split(line, " ")
	cnt := 0
	for _, w := range words {
		if w != "" {
			cnt++
		}
	}

	fmt.Print(cnt)
}
