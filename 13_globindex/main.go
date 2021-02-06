package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml", "dsindex.gohtml"))
}

func main() {

	xs := []string{"zero", "one", "two", "three", "four", "five"}

	data := struct {
		Word  []string
		Lname string
	}{
		xs,
		"last",
	}

	err := tpl.ExecuteTemplate(os.Stdout, "index.gohtml", xs)
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "dsindex.gohtml", data)
	if err != nil {
		log.Fatalln(err)
	}
}
