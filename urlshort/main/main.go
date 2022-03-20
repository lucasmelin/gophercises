package main

import (
	"fmt"
	"net/http"

	"github.com/lucasmelin/gophercises/urlshort"
)

func main() {
	mux := defaultMux()

	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}

	mapHandler := urlshort.MapHandler(pathsToUrls, mux)

	// YAML requires spaces for indentation
	yaml := `
    - path: /urlshort
      url: https://github.com/lucasmelin/gophercises/urlshort
    `

	yamlHandler, err := urlshort.YAMLHandler([]byte(yaml), mapHandler)
	if err != nil {
		panic(err)
	}
	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", yamlHandler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(writer http.ResponseWriter, requestPtr *http.Request) {
	fmt.Fprintln(writer, "Hello, world!")
}
