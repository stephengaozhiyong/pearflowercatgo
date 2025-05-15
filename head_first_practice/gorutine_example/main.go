package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

type Page struct {
	url  string
	size int
}

func responseSize(url string, ch chan Page) {
	fmt.Println("Getting:", url)
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	ch <- Page{url: url, size: len(body)}
}

func main() {
	c := make(chan Page)
	go responseSize("https://example.com", c)
	go responseSize("https://golang.org", c)
	go responseSize("https://golang.org/doc", c)
	for i := range c {
		fmt.Println(i)
	}
}
