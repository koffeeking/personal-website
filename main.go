package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

const port string = ":8081"

// reads HTML templates from the specified directory
func compileHelloTemplates() *template.Template {
	parsedTemplates, err := template.ParseFiles(
		filepath.Join("templates", "index.html"),
		filepath.Join("templates", "hello.html"),
	)
	if err != nil {
		log.Fatalf("Failed to parse tempaltes: %v", err)
	}
	return parsedTemplates
}

// returns a simple Hello World response.
func helloHandler(tmpl *template.Template) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		if request.Header.Get("HX-Request") == "true" {
			// render partial hello.html for HTMX request
			tmpl.ExecuteTemplate(writer, "hello.html", nil)
		} else {
			// render full index.html when accessing root
			tmpl.ExecuteTemplate(writer, "index.html", nil)
			tmpl.ExecuteTemplate(writer, "sayHi.html", nil)
		}
		log.Println(request.Header)
	}
}

func main() {
	helloTemplate := compileHelloTemplates()
	http.HandleFunc("/", helloHandler(helloTemplate))

	fmt.Println("Server started at ", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		fmt.Println("Failed to start server:", err)
	}
}
