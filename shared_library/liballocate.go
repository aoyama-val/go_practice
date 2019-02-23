package main

import (
	"C"
	"log"
)

//export allocate
func allocate(size int) {
	buf := make([]int, size)
	for j := 0; j < size; j++ {
		buf[j] = j
	}
	//fmt.Printf("allocate in Go\n")
}

func init() {
	log.Println("Loaded!!")
}

func main() {
}
