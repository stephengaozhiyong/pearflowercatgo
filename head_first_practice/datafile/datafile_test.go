package datafile

import (
	"fmt"
	"log"
	"testing"
)

func TestGetFloats(t *testing.T) {
	numbers, err := GetFloats("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	for i, n := range numbers {
		fmt.Println(i, n)
	}

}
