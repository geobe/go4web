package main

import (
	"fmt"
	"github.com/geobe/go4j/poi"
	"html/template"
	"net/http"
	"os"
	"strconv"
)

const html = `<!DOCTYPE html>
<html>
<head>
<meta content="text/html;
charset=windows-1252" http-equiv="content-type">
<title>City Info</title>
</head>
<body>
Stadt: %s<br>
Einwohner: %d<br>
Lat, Lon: %f, %f<br>
</p>
</body>
</html>
`

var cityTpl *template.Template

// Parse die angegebenen Template-Files in ein Template
// templatedir	Verzeichnis mit den Template-Files
// fn...	ein oder mehrere Filenamen (ohne .html)
func prep4(templatedir string, fn ...string) (t *template.Template) {
	var files []string
	for _, file := range fn {
		files = append(files, fmt.Sprintf("%s/%s.html", templatedir, file))
	}
	t = template.Must(template.ParseFiles(files...))
	return
}

func parseCity(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	index, err := strconv.Atoi(r.PostFormValue("city"))
	if err != nil {
		fmt.Fprintf(w, "Error, index %s is no number!", r.PostFormValue("city"))
		return
	}
	city := poi.GermanCities[index]
	lat, lon := city.LatLon()
	fmt.Fprintf(w, html, city.Name(), city.Inhabitants(), lat, lon)
}

func cityForm(w http.ResponseWriter, r *http.Request) {
	citynames := make([]string, len(poi.GermanCities))
	for i, c := range poi.GermanCities {
		citynames[i] = c.Name()
	}
	if cityTpl != nil {
		cityTpl.ExecuteTemplate(w, "cityform", citynames)
	}
}

func main() {
	mux := http.NewServeMux()
	pwd, _ := os.Getwd()
	tpl := pwd + "/src/github.com/geobe/go4web/webmain1/tpl"
	cityTpl = prep4(tpl, "CityForm")
	mux.HandleFunc("/eval", parseCity)
	mux.HandleFunc("/", cityForm)
	server := &http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: mux,
	}
	server.ListenAndServe()
}
