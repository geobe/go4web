package main

import (
	"fmt"
	"net/http"
	"os"
)

const html = `<!DOCTYPE html>
<html>
<head>
<meta content="text/html;
charset=windows-1252" http-equiv="content-type">
<title>Echo page</title>
</head>
<body>
<p>Commit URL: %s <br>
Nickname: %s<br>
Slogan: %s<br>
Importance: %s<br>
</p>
</body>
</html>
`

func parse(writer http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Fprintf(writer, html, r.URL.Path[:], r.PostFormValue("nickname"),
		r.PostFormValue("slogan"), r.PostFormValue("importance"))
}

func main() {
	mux := http.NewServeMux()
	pwd, _ := os.Getwd()
	dir := http.Dir(pwd + "/src/github.com/geobe/go4web/webmain1/pub")
	files := http.FileServer(dir)
	mux.Handle("/static/", http.StripPrefix("/static/", files))
	mux.HandleFunc("/", parse)
	server := &http.Server{
		Addr:    "127.0.0.1:8110",
		Handler: mux,
	}
	server.ListenAndServe()
}
