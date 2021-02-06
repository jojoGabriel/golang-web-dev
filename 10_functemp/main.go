package main

import (
	"html/template"
	"log"
	"os"
	"strings"
)

type writer struct {
	Name string
	Book string
}

var funcMap = template.FuncMap{
	"uc": strings.ToUpper,
}

func main() {

	tmplt := template.Must(template.New("").Funcs(funcMap).ParseFiles("funcmap.gohtml"))

	m := writer{Name: "Matt", Book: "King"}
	j := writer{Name: "John", Book: "God"}

	w := []writer{m, j}

	data := struct {
		Writer []writer
	}{
		w,
	}

	err := tmplt.ExecuteTemplate(os.Stdout, "funcmap.gohtml", data)
	if err != nil {
		log.Fatalln(err)
	}
}
