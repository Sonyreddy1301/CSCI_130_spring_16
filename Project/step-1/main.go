package main

import (
	"html/template"
	"log"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	tpl, err := template.ParseFiles("application.html")
	if err != nil {

		log.Fatalln(err) //timestamp

	}
	err = tpl.Execute(w, nil)
}
func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)

}
