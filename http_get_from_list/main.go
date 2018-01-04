package main

import (
	"fmt"
)

func main() {
	lines := readLines("url.txt")
	for _, line := range lines {
		fmt.Println(line)
		get(line)
	}
}
