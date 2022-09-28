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
	fmt.Println(lines)
}
