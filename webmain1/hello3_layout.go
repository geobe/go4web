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

	// hole einen DefaultServeMux
	mux := http.NewServeMux()
	// finde Working directory = GOPATH
	pwd, _ := os.Getwd()
	// und hänge den ganzen Pfad zu den statischen Files dahinter
	dir := http.Dir(pwd + "/src/github.com/geobe/go4j/webmain1/pub")
	files := http.FileServer(dir)
	// unter der URL /static/ werden files bereitgestellt,
	// Präfix /static/ wird abgeschnitten
	mux.Handle("/static/", http.StripPrefix("/static/", files))
	// index
	mux.HandleFunc("/", parse)
	// konfiguriere server
	server := &http.Server{
		Addr:    "0.0.0.0:8110",
		Handler: mux,
	}
	// und starte ihn
	server.ListenAndServe()
}
