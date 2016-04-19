package myapp

import (
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
	"google.golang.org/cloud/storage"
	"io"
	"mime/multipart"
	"net/http"
	"strings"
)

const gcsBucket = "csci_130_sony"

type Entity struct {
	Value string
}

func init() {
	http.HandleFunc("/", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
	context := appengine.NewContext(r)

	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	html := `
	    <form method="POST" enctype="multipart/form-data">
		<input type="file" name="fileToUpload">
		<input type="submit">
	    </form>
	`

	if r.Method == "POST" {

		myMultiPartFile, myHeader, err := r.FormFile("fileToUpload")

		if err != nil {
			log.Errorf(context, "error handler r.FormFile: ", err)
			http.Error(w, "we were unable to upload your file \n", 500)
			return
		}

		defer myMultiPartFile.Close()

		fname, err := uploadFile(r, myMultiPartFile, myHeader)
		if err != nil {
			log.Errorf(context, "error handler uploadFile:", err)
			http.Error(w, "we were unable to accept your file \n"+err.Error(), 415)
			return
		}

		fnames, err := putCookie(w, r, fname)
		if err != nil {
			log.Errorf(context, "error handler putCookie:", err)
			http.Error(w, "we were unable to accept your file \n"+err.Error(), 415)
			return
		}

		html += `<h1>Files</h1>`

		for k, _ := range fnames {
			html += `<h3>` + k + `</h3>`
		}
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, html)
}

func putCookie(w http.ResponseWriter, r *http.Request, fname string) (map[string]bool, error) {

	mss := make(map[string]bool)
	cookie, _ := r.Cookie("file-names")
	if cookie != nil {
		bs, err := base64.URLEncoding.DecodeString(cookie.Value)
		if err != nil {
			return nil, fmt.Errorf("error handler base 54.URLEnoding.DecodeString:: %s", err)
		}
		err = json.Unmarshal(bs, &mss)
		if err != nil {
			return nil, fmt.Errorf("error handler json.Unmarshal: %s)", err)
		}
	}

	mss[fname] = true
	bs, err := json.Marshal(mss)
	if err != nil {
		return mss, fmt.Errorf("error putCookie json.Marshal: ", err)
	}

	b64 := base64.URLEncoding.EncodeToString(bs)

	// FYI
	ctx := appengine.NewContext(r)
	log.Infof(ctx, "COOKIE JSON: %s", string(bs))

	http.SetCookie(w, &http.Cookie{
		Name:  "file-names",
		Value: b64,
	})
	return mss, nil
}

func uploadFile(req *http.Request, mpf multipart.File, hdr *multipart.FileHeader) (string, error) {

	ext, err := fileFilter(req, hdr)
	if err != nil {
		return "", err
	}
	name := getSha(mpf) + `.` + ext
	mpf.Seek(0, 0)

	ctx := appengine.NewContext(req)
	return name, putFile(ctx, name, mpf)
}

func fileFilter(req *http.Request, hdr *multipart.FileHeader) (string, error) {

	ext := hdr.Filename[strings.LastIndex(hdr.Filename, ".")+1:]
	ctx := appengine.NewContext(req)
	log.Infof(ctx, "FILE EXTENSION: %s", ext)

	switch ext {
	case "jpg", "jpeg", "txt", "md":
		return ext, nil
	}
	return ext, fmt.Errorf("We do not allow files of type %s. We only allow jpg, jpeg, txt, md extensions.", ext)
}

func getSha(src multipart.File) string {
	h := sha1.New()
	io.Copy(h, src)
	return fmt.Sprintf("%x", h.Sum(nil))
}

func putFile(ctx context.Context, name string, rdr io.Reader) error {

	client, err := storage.NewClient(ctx)
	if err != nil {
		return err
	}
	defer client.Close()

	writer := client.Bucket(gcsBucket).Object(name).NewWriter(ctx)
	writer.ACL = []storage.ACLRule{
		{storage.AllUsers, storage.RoleReader},
	}
	io.Copy(writer, rdr)
	return writer.Close()
}
