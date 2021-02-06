package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	// parseExecStdout()
	// parseExecFile()
	// executeTemplates()
	parseGlob()

}

func parseExecStdout() {
	tmplt, err := template.ParseFiles("hello.html")
	if err != nil {
		log.Fatalln(err)
	}

	err = tmplt.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}
}

func parseExecFile() {
	tmpl, err := template.ParseFiles("hello.gohtml")
	if err != nil {
		log.Fatalln("unable to parse file", err)
	}

	fil, err := os.Create("index.html")
	if err != nil {
		log.Fatalln("error creating file", err)
	}
	defer fil.Close()

	err = tmpl.Execute(fil, nil)
	if err != nil {
		log.Fatalln()
	}

}

func executeTemplates() {
	pt, err := template.ParseFiles("one.gohtml")
	if err != nil {
		log.Fatalln(err)
	}

	err = pt.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}

	pt, err = pt.ParseFiles("two.gohtml", "three.gohtml")
	if err != nil {
		log.Fatalln(err)
	}

	err = pt.ExecuteTemplate(os.Stdout, "three.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}

	err = pt.ExecuteTemplate(os.Stdout, "two.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}

	err = pt.ExecuteTemplate(os.Stdout, "one.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}

	err = pt.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}
}

func parseGlob() {
	pt, err := template.ParseGlob("templates/*")
	if err != nil {
		log.Fatalln(err)
	}

	err = pt.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}

	pt, err = pt.ParseFiles("two.gohtml", "three.gohtml")
	if err != nil {
		log.Fatalln(err)
	}

	err = pt.ExecuteTemplate(os.Stdout, "three.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}

	err = pt.ExecuteTemplate(os.Stdout, "two.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}

	err = pt.ExecuteTemplate(os.Stdout, "one.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}
}
