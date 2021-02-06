package main

import (
	"log"
	"os"
	"text/template"
)

var tmplt *template.Template

type teamCaptain struct {
	Captain string
	Team    string
}

func init() {
	tmplt = template.Must(
		template.ParseFiles(
			"range.gohtml",
			"indexValue.gohtml",
			"struct.gohtml",
			"structProp.gohtml",
			"structSlices.gohtml",
			"structSlicesStruct.gohtml"))
}

func main() {

	players := []string{"Jordan", "Magic", "Bird"}
	captain := map[string]string{
		"Jordan": "Bulls",
		"Magic":  "Lakers",
		"Bird":   "Celtics",
	}

	curry := teamCaptain{
		Captain: "Curry",
		Team:    "Warriors",
	}

	irving := teamCaptain{
		Captain: "Irving",
		Team:    "Brooklyn",
	}

	teamCaptains := []teamCaptain{curry, irving}

	bestHandlers := struct {
		Guard []teamCaptain
	}{
		teamCaptains,
	}

	err := tmplt.ExecuteTemplate(os.Stdout, "range.gohtml", players)
	if err != nil {
		log.Fatalln(err)
	}

	err = tmplt.ExecuteTemplate(os.Stdout, "indexValue.gohtml", players)
	if err != nil {
		log.Fatalln(err)
	}

	err = tmplt.ExecuteTemplate(os.Stdout, "range.gohtml", captain)
	if err != nil {
		log.Fatalln(err)
	}

	err = tmplt.ExecuteTemplate(os.Stdout, "indexValue.gohtml", captain)
	if err != nil {
		log.Fatalln(err)
	}

	err = tmplt.ExecuteTemplate(os.Stdout, "struct.gohtml", curry)
	if err != nil {
		log.Fatalln(err)
	}

	err = tmplt.ExecuteTemplate(os.Stdout, "structProp.gohtml", curry)
	if err != nil {
		log.Fatalln(err)
	}

	err = tmplt.ExecuteTemplate(os.Stdout, "structSlices.gohtml", teamCaptains)
	if err != nil {
		log.Fatalln(err)
	}

	err = tmplt.ExecuteTemplate(os.Stdout, "structSlicesStruct.gohtml", bestHandlers)
	if err != nil {
		log.Fatalln(err)
	}
}
