package main

import(
    "fmt"
    "net/http"


)

func formvalue( w http.ResponseWriter, r *http.Request){
    name := r.FormValue("name")
    fmt.Fprintf(w, "Hi form value is working %s", name)
}

func main(){
  http.HandleFunc("/", formvalue)
	http.ListenAndServe(":8080", nil)

}
