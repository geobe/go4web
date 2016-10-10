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
func prepWithFuncs(funcmap template.FuncMap, templatedir string, fn ...string) (t *template.Template) {
	var files []string
	for _, file := range fn {
		files = append(files, fmt.Sprintf("%s/%s.html", templatedir, file))
	}
	// Functions bei neuem Template registrieren
	t = template.New("cf")
	t.Funcs(funcmap)
	template.Must(t.ParseFiles(files...))
	return
}

func parseCity5(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	index, err := strconv.Atoi(r.PostFormValue("city"))
	if err != nil {
		fmt.Fprintf(w, "Error, index \"%s\" is no number!", r.PostFormValue("city"))
		return
	}
	city := poi.GermanCities[index]
	lat, lon := city.LatLon()
	fmt.Fprintf(w, html, city.Name(), city.Inhabitants(), lat, lon)
}

func cityForm5(w http.ResponseWriter, r *http.Request) {
	if cityTpl != nil {
		// Array von Objekten -> Template
		cityTpl.ExecuteTemplate(w, "cityform",
			poi.GermanCities)
	}
}

func main() {
	mux := http.NewServeMux()
	pwd, _ := os.Getwd()
	tpl := pwd + "/src/github.com/geobe/go4web/webmain1/tpl"
	// definition der Funktions-Map
	funcMap := template.FuncMap{
		"name": (poi.City).Name}
	cityTpl = prepWithFuncs(funcMap, tpl, "CityFormMethod")
	mux.HandleFunc("/eval", parseCity5)
	mux.HandleFunc("/", cityForm5)
	server := &http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: mux,
	}
	server.ListenAndServe()
}
