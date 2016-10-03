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
<title>Generated index page</title>
</head>
<body>
<p>Hallo Welt,<br>
du hast die URL %s aufgerufen.<br>
<a href="static/HelloStatics.html">
Geh doch mal zu einer statischen Seite</a>.<br>
</p>
</body>
</html>
`

func index(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, html, request.URL.Path[1:])
}

func err(writer http.ResponseWriter,
	request *http.Request) {
	fmt.Fprintf(writer, "Error with path %s",
		request.URL.Path[1:])
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
	mux.HandleFunc("/index", index)
	// error
	mux.HandleFunc("/err", err)
	// konfiguriere server
	server := &http.Server{
		Addr:    "0.0.0.0:8090",
		Handler: mux,
	}
	// und starte ihn
	server.ListenAndServe()
}
