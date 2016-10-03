package main

import (
	"fmt"
	"net/http"
)

func hello(writer http.ResponseWriter,
	request *http.Request) {
	fmt.Fprintf(writer,
		"Hello from Go, %s!",
		request.URL.Path[1:])
}

func main() {
	http.HandleFunc("/", hello)
	http.ListenAndServe(":8080", nil)
}
