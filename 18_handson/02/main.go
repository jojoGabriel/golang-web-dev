package main

import (
	"log"
	"os"
	"text/template"
)

type hotel struct {
	Name, Address, City, Zip, Region string
}

var hotels []hotel

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {

	hotels = []hotel{
		hotel{
			Name:    "ABC",
			Address: "ABC Street",
			City:    "ABC City",
			Zip:     "ABC123",
			Region:  "Northern",
		},
		hotel{
			Name:    "DEF",
			Address: "DEF Street",
			City:    "DEF City",
			Zip:     "DEF456",
			Region:  "Southern",
		},
	}

	err := tpl.Execute(os.Stdout, hotels)
	if err != nil {
		log.Fatalln(err)
	}
}
