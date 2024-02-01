package main

import (
	"fmt"
	"net/http"
	"url-shortener/utils"
)

func main() {
	mux := defaultMux()

	// Build the MapHandler using the mux as the fallback
	pathsToUrls := map[string]string{
		"/shortener-godoc": "https://godoc.org/github.com/gophercises/shortener",
		"/yaml-godoc":      "https://godoc.org/gopkg.in/yaml.v2",
	}
	mapHandler := utils.MapHandler(pathsToUrls, mux)

	// Build the YAMLHandler using the mapHandler as the
	// fallback
	yaml := `
  - path: /shortener
    url: https://github.com/gophercises/shortener
  - path: /shortener-final
    url: https://github.com/gophercises/shortener/tree/solution
  `
	yamlHandler, err := utils.YAMLHandler([]byte(yaml), mapHandler)
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

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}
