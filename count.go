package main

import (
	"fmt"
	"log"

	"github.com/mlarikov/go_headfirst_datafile"
)

func main() {
	lines, err := go_headfirst_datafile.GetStrings("votes.txt")
	if err != nil {
		log.Fatal(err)
	}
	counts := make(map[string]int)
	for _, line := range lines {
		counts[line]++
	}
	fmt.Println(counts)
}
