package main

import (
	"fmt"
	"net/http"
)

func formvalue(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, r.FormValue("name"))
}

func main() {
	http.HandleFunc("/", formvalue)
	http.ListenAndServe(":8080", nil)

}
