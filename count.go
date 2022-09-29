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
	var names []string
	var counts []int
	for _, line := range lines {
		matched := false
		for i, name := range names {
			if name == line {
				counts[i]++
				matched = true
			}
		}
		if !matched {
			names = append(names, line)
			counts = append(counts, 1)
		}
	}
	for i, name := range names {
		fmt.Printf("%s: %d\n", name, counts[i])
	}
}
