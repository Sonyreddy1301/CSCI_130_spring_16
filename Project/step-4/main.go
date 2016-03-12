package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/nu7hatch/gouuid"
	"net/http"
  "io"
)

type User struct {
	Name string
	Age  string
}

func cookieid(w http.ResponseWriter, r *http.Request) {
		name := r.FormValue("name")
		age := r.FormValue("age")
		data := foo(name, age)


	cookie, err := r.Cookie("sessio-info")

	if err != nil {

		id, _ := uuid.NewV4()
		cookie = &http.Cookie{
			Name:  "session-info",
			Value: id.String() + "|" + name + age + "|" + data,
			HttpOnly: true,
		}
		http.SetCookie(w, cookie)
    io.WriteString(w, `<!DOCTYPE html>
      <html>
      	<head>
      		<title>Project-4</title>
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
func foo(name string, age string) string {
	user := User{
		Name: name,
		Age:  age,
	}

	bs, err := json.Marshal(user)
	if err != nil {
		fmt.Printf("error: ", err)
	}
	str := base64.URLEncoding.EncodeToString(bs)
	return str
}
func main() {

	http.HandleFunc("/", cookieid)
	http.ListenAndServe(":8080", nil)
}
