package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	max := 1000
	var err error

	if len(os.Args) >= 2 {
		max, err = strconv.Atoi(os.Args[1])
		if err != nil {
			max = 1000
		}
	}

	var flags = make([]bool, max)

	for i := range flags {
		flags[i] = true
	}

	for n := 2; n < max; n++ {
		if !flags[n] {
			continue
		}
		for x := n + n; x < max; x += n {
			flags[x] = false
		}
	}

	for i := 2; i < max; i++ {
		if flags[i] {
			fmt.Printf("%d ", i)
		}
	}

	fmt.Print("\n")
}
