package main

import (
	"fmt"
	"io"
	"net/http"
)

func formvalue(w http.ResponseWriter, r *http.Request) {
	key := "text file"
	file, hdr, err := r.FormValue(key)
	fmt.Println(file, hdr, err)
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `<form method="POST" enctype="multipart/form-data">
      <input type="file" name="text file">
      <input type="submit">
    </form>`)
}
func main() {
	http.HandleFunc("/", formvalue)
	http.ListenAndServe(":8080", nil)

}
