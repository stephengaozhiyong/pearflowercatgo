package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

func Walk(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}
	Walk(t.Left, ch)
	ch <- t.Value
	Walk(t.Right, ch)
}

func Same(t1, t2 *tree.Tree) bool {
	c1 := make(chan int)
	go func() {
		defer close(c1)
		Walk(t1, c1)
	}()

	c2 := make(chan int)
	go func() {
		defer close(c2)
		Walk(t2, c2)
	}()

	for {
		v1, ok1 := <-c1
		v2, ok2 := <-c2

		if !ok1 || !ok2 {
			if ok1 != ok2 {
				return false
			}
			break
		}
		if ok1 && ok2 && v1 != v2 {
			return false
		}
	}
	return true

}

func main() {
	if Same(tree.New(3), tree.New(3)) {
		fmt.Println("ok")
	} else {
		fmt.Println("failed")
	}
}
