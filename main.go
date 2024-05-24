package main

import (
	"go_web_app/entities"
	"net/http"
	"text/template"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	http.HandleFunc("/", index)
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		return
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	products := []entities.Product{
		{"Xícara", "Azul com desenho", 35.94, 40},
		{"Sapato", "Couro", 99.90, 100},
		{"Tenis", "Confortável", 89, 3},
	}

	err := temp.ExecuteTemplate(w, "Index", products)
	if err != nil {
		return
	}
}
