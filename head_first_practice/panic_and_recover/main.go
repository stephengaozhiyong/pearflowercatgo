package main

import (
	"fmt"
	"os"
	"path"
	"strings"
)

func reportPanic() {
	p := recover()
	if p == nil {
		return
	}
	e, ok := p.(error)
	if ok {
		fmt.Println(e.Error())
	} else {
		panic(p)
	}
}

func ListDir(file_path string, depth int) {
	dir_entry, err := os.ReadDir(file_path)
	if err != nil {
		panic(err)
	}
	for _, dr := range dir_entry {
		if dr.IsDir() {
			fmt.Println(strings.Repeat("\t", depth), "Dir:", dr.Name())
			ListDir(path.Join(file_path, dr.Name()), depth+1)
		} else {
			fmt.Println(strings.Repeat("\t", depth), "File:", dr.Name())
		}
	}
}

func main() {
	defer reportPanic()
	panic("sdfds")
	ListDir(os.Args[1], 0)
}
