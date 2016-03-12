package main

import (
	"fmt"
	"github.com/nu7hatch/gouuid"
	"net/http"
)

func cookieid(w http.ResponseWriter, r *http.Request) {

	cookie, err := r.Cookie("sessio-id")

	if err != nil {
	//	fmt.Fprint(w, "Creating the cookie...\n"
		id, _ := uuid.NewV4()
		cookie = &http.Cookie{
			Name:  "session-id",
			Value: id.String(),
			// Secure: true,
			HttpOnly: true,
		}
		http.SetCookie(w, cookie)
	}
	fmt.Fprint(w, "Name: ", cookie.Name, "\nValue: ", cookie.Value, "\nHttpOnly: ", cookie.HttpOnly)

}
func main() {

	http.HandleFunc("/", cookieid)
	http.ListenAndServe(":8080", nil)
}
