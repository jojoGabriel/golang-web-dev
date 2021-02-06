package main

import (
	"log"
	"os"
	"text/template"
)

var tmplt *template.Template

func init() {
	tmplt = template.Must(template.ParseFiles("template.gohtml"))
}

func main() {
	err := tmplt.ExecuteTemplate(os.Stdout, "template.gohtml", `you reached 50`)
	if err != nil {
		log.Fatalln(err)
	}
}
