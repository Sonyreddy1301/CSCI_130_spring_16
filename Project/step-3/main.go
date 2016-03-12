package main

import (

	"github.com/nu7hatch/gouuid"
	"net/http"
  "fmt"
  "io"
)

func cookieid(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	age := r.FormValue("age")
	cookie, err := r.Cookie("sessio-info")

	if err != nil {

		id, _ := uuid.NewV4()
		cookie = &http.Cookie{
			Name:  "session-info",
			Value: id.String() + "|" + name + age,
			HttpOnly: true,
		}
		http.SetCookie(w, cookie)
    io.WriteString(w, `<!DOCTYPE html>
      <html>
      	<head>
      		<title>Project-3</title>
      	</head>
      	<body>
      		<form method="POST">
      			Name: <input type="text" name="name"><br>
      			Age:  <input type="text" name="age"><br>
      			<input type="submit">
      		</form>
      	</body>
      </html>
	`)
	}

fmt.Fprint(w,"Value: ", cookie.Value )
}
func main() {

	http.HandleFunc("/", cookieid)
	http.ListenAndServe(":8080", nil)
}
