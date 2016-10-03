package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"time"
)

var template1 *template.Template

func idx(writer http.ResponseWriter, request *http.Request) {
	// aktuelle Zeit einschließlich Datum
	now := time.Now()
	// eine struct zur Datenübergabe an das Template
	var dt struct {
		Title, Time, Date string
	}
	dt.Time = now.Format("15:04:05.9999")
	dt.Date = now.Format("1.2.2006")
	if len(request.URL.Path) > 1 {
		dt.Title = request.URL.Path[1:]
	} else {
		dt.Title = "Hallo Welt!"
	}
	if template1 != nil {
		template1.ExecuteTemplate(writer, "datetime", dt)
	}
}

func err1(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Error with path %s",
		request.URL.Path[1:])
}

func prep(templatedir string, fn ...string) (t *template.Template) {
	var files []string
	for _, file := range fn {
		files = append(files, fmt.Sprintf("%s/%s.html", templatedir, file))
	}
	t = template.Must(template.ParseFiles(files...))
	return
}

func main() {

	// hole einen DefaultServeMux
	mux := http.NewServeMux()
	// finde Working directory = GOPATH
	pwd, _ := os.Getwd()
	// und hänge den ganzen Pfad zu den statischen Files dahinter
	tpl := pwd + "/src/github.com/geobe/go4j/webmain1/tpl"
	template1 = prep(tpl, "HelloTime")
	// und hänge den ganzen Pfad zu den statischen Files dahinter
	dir := http.Dir(pwd + "/src/github.com/geobe/go4j/webmain1/pub")
	files := http.FileServer(dir)
	// unter der URL /static/ werden files bereitgestellt,
	// Präfix /static/ wird abgeschnitten
	mux.Handle("/static/", http.StripPrefix("/static/", files))
	// index
	mux.HandleFunc("/", idx)
	// error
	mux.HandleFunc("/err", err1)
	// konfiguriere server
	server := &http.Server{
		Addr:    "0.0.0.0:8100",
		Handler: mux,
	}
	// und starte ihn
	server.ListenAndServe()
}
