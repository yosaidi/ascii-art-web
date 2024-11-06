package main

import (
	"fmt"
	"html/template"
	"net/http"

	ascii "ascii/AsciiFunctions"
)

type TemplateData struct {
	Result string
}

var tmpl = template.Must(template.ParseFiles("templates/index.html"))

func main() {
	http.HandleFunc("/", PagesHandler)

	fmt.Println("Server started at localhost:8080!")
	http.ListenAndServe(":8080", nil)
}

func PagesHandler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		MainPage(w, r)
	case "/ascii-art":
		AsciiPage(w, r)
	default:
		http.NotFound(w, r)
	}
}

func MainPage(w http.ResponseWriter, r *http.Request) {
	tmpl.Execute(w, nil)
}

func AsciiPage(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		text := r.FormValue("input")
		banner := r.FormValue("banner")

		// Check for input validation
		if err := ascii.ErrorHandler(text, banner); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		result, err := ascii.Generate(text, banner)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		data := TemplateData{Result: result}
		tmpl.Execute(w, data)

	} else {
		http.Error(w, "Bad Request:  Only POST method allowed", http.StatusBadRequest)
	}
}
