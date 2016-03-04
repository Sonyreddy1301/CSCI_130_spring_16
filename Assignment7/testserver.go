package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)
}

func index(res http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/" {
		http.NotFound(res, req)
		return
	}
	io.WriteString(res, "I know where I'm going and I know the truth, and I don't have to be what you want me to be. I'm free to be what I want.--Mohammad Ali")
}
