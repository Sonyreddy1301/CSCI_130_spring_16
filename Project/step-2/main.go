package main

import (
	"github.com/nu7hatch/gouuid"
	"html/template"
	"log"
	"net/http"
)

func cookieid(w http.ResponseWriter, r *http.Request) {
	tpl, err := template.ParseFiles("index.html")
	if err != nil {
		log.Fatalln(err)
	}

	cookie, err := r.Cookie("sessio-info")

	if err != nil {

		id, _ := uuid.NewV4()
		cookie = &http.Cookie{
			Name:  "session-info",
			Value: id.String(),
			//Secure:   true,
			HttpOnly: true,
		}
		http.SetCookie(w, cookie)
	}
	//fmt.Fprint(w, "Name: ", cookie.Name, "\nValue: ", cookie.Value, "\nHttpOnly: ", cookie.HttpOnly)
	err = tpl.Execute(w, nil)

}
func main() {

	http.HandleFunc("/", cookieid)
	http.ListenAndServe(":8080", nil)
}
