
package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl, _ = template.ParseFiles("text.html")
}

func name(w http.ResponseWriter, r *http.Request) {
	tpl.Execute(w, nil)
	fmt.Fprintf(w, "%v", r.FormValue("n"))
}

func main() {
	http.HandleFunc("/", name)
	http.ListenAndServe(":8080", nil)
}
