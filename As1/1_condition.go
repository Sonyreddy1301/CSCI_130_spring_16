package main

import (
	"log"
	"os"
	"text/template"
)

type degree struct {
	Name   string
	Stream string
}
type expertise struct {
	degree
	isEligible bool
}

func main() {
	exp := expertise{
		degree: degree{
			Name:   "Sony",
			Stream: "CSE",
		},
		isEligible: true,
	}
	tpl, err := template.ParseFiles("tpl.html")

	if err != nil {
		log.Fatalln(err)
	}
	err = tpl.Execute(os.Stdout, p1)
	if err != nil {
		log.Fatalln(err)
	}
}
