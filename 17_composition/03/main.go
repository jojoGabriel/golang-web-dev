package main

import (
	"html/template"
	"log"
	"os"
)

type course struct {
	Number, Name, Units string
}

type semester struct {
	Term    string
	Courses []course
}

type year struct {
	Fall, Spring, Summer semester
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	yr := year{
		Fall: semester{
			Term: "Fall",
			Courses: []course{
				course{"CS01", "Go Programming", "4"},
				course{"CS02", "Go Web Programming", "4"},
				course{"CS03", "Go Mobile", "4"},
			},
		},
		Spring: semester{
			Term: "Spring",
			Courses: []course{
				course{"ACS01", "Advanced Go Programming", "4"},
				course{"ACS02", "Advanced Go Web Programming", "4"},
				course{"ACS03", "Advanced Go Mobile", "4"},
			},
		},
	}

	err := tpl.Execute(os.Stdout, yr)
	if err != nil {
		log.Fatalln(err)
	}
}
