package main

import (
	"fmt"
	"log"

	"headfirstgo/datafile"
)

func main() {
	lines, err := datafile.GetStrings("/home/sgao/go/src/headfirstgo/main/votes.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(lines)
	var names []string
	var counts []int
	for _, line := range lines {
		matched := false
		for i, name := range names {
			if name == line {
				matched = true
				counts[i]++
			}
		}
		if !matched {
			names = append(names, line)
			counts = append(counts, 1)
		}
	}
	for i, name := range names {
		fmt.Printf("%s: %d\n", name, i)
	}
}
