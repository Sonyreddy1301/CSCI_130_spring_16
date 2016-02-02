package main

import (
	"html/template"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", surfWebPage)
	http.Handle("/files/", http.StripPrefix("/files/", http.FileServer(http.Dir("files"))))
	http.ListenAndServe(":9000", nil)
}

func surfWebPage(res http.ResponseWriter, req *http.Request) {

	var err error

	surferPage := template.New("Surfer Page.html")
	surferPage, err = surferPage.ParseFiles("Surfer Page.html")

	if err != nil {
		log.Fatalln(err)
	}

	err = practiceWebPage.Execute(res, nil)

	if err != nil {
		log.Fatalln(err)
	}
}
