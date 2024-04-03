package main

import (
	"fmt"
	"headfirstgo/prose"
)

func main() {
	a := []string{"我", "叫", "高", "志", "勇"}
	res := prose.JoinWithCommas(a)
	fmt.Println(res)

}
