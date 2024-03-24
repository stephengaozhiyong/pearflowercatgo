package main

import (
	"datafile"
	"fmt"
	"log"
)

func main() {
	nums, err := datafile.GetFloats("/home/sgao/goPractice/temp/data.txt")
	if err != nil {
		log.Fatal(err)
	}
	var sum float64 = 0
	for _, n := range nums {
		sum += n
	}
	fmt.Println(sum)
	count := float64(len(nums))
	avg := sum / count
	fmt.Printf("Average Value: %0.2f\n", avg)
}
