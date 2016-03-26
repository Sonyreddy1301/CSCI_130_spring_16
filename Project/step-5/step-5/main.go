package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/nu7hatch/gouuid"
	//"html/template"
	"io"
	//"log"
	"net/http"
)

type User struct {
	Name string
	Age  string
}

func cookieid(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	age := r.FormValue("age")
	data := foo(name, age)
	code := getCode(data)

	cookie, err := r.Cookie("sessio-info")

	if err != nil {

		id, _ := uuid.NewV4()
		cookie = &http.Cookie{
			Name:     "session-info",
			Value:     id.String() + "|" + name + "|" + age + "|" + data + "|" + code,
			HttpOnly: true,
		}
		http.SetCookie(w, cookie)
		io.WriteString(w, `<!DOCTYPE html>
      <html>
      	<head>
      		<title>Project-5</title>
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
	fmt.Fprint(w, "Value: ", cookie.Value)
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
func getCode(data string) string {
	h := hmac.New(sha256.New, []byte("key"))
	io.WriteString(h, data)
	return fmt.Sprintf("%x", h.Sum(nil))
}
func main() {

	http.HandleFunc("/", cookieid)
	http.ListenAndServe(":8080", nil)
}
