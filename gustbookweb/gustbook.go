package main

import (
	"fmt"
	"headfirstgo/datafile"
	"log"
	"net/http"
	"os"
	"text/template"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func viewHandler(writer http.ResponseWriter, request *http.Request) {
	signatures, err := datafile.GetStrings("signatures.txt")
	check(err)
	fmt.Printf("%#v\n", signatures)
	html, err := template.ParseFiles("view.html")
	check(err)
	err = html.Execute(writer, nil)
	check(err)
}

func exeTemplate(text string, data any) {
	tmpl, err := template.New("test").Parse(text)
	check(err)
	err = tmpl.Execute(os.Stdout, data)
	check(err)
}

type Guestbook struct {
	SignatureCount int
	Signatures     []string
}

func main() {
	http.HandleFunc("/gustbook", viewHandler)
	err := http.ListenAndServe("localhost:8080", nil)
	log.Fatal(err)
}
