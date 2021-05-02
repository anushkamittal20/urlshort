package main

import (
	"fmt"
	"net/http"

	"github.com/anushkamittal20/urlshort"
)

func main() {
	mux := defaultMux()

	// Build the MapHandler using the mux as the fallback
	pathsToUrls := map[string]string{
		"/w3":        "https://www.w3schools.com/html/",
		"/html-wiki": "https://en.wikipedia.org/wiki/HTML",
	}
	mapHandler := urlshort.MapHandler(pathsToUrls, mux)

	// Build the YAMLHandler using the mapHandler as the
	// fallback
	yaml := `
- path: /urlshort
  url: https://github.com/anushkamittal20/urlshort
- path: /link
  url: https://github.com/anushkamittal20/link
- path: /urlshort-final
  url: https://github.com/gophercises/urlshort/tree/solution
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

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world! \n to redirect to a new page try entering http://localhost:8080/w3 or \n http://localhost:8080/html-wiki or \n to redirect to my go projects on github \n typehttp://localhost:8080/link or http://localhost:8080/urlshort \n Thank you")
}
