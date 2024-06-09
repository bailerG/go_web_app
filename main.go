package main

import (
	"go_web_app/entities"
	"go_web_app/repository"
	"net/http"
	"text/template"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	http.HandleFunc("/", index)
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		return
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	db := database.GetDB()

	result, err := db.Query("select * from produtos")
	if err != nil {
		panic(err.Error())
	}

	p := entities.Product{}
	var products []entities.Product

	for result.Next() {
		var id, quantity int
		var name, description string
		var price float64

		err := result.Scan(&id, &name, &description, &price, &quantity)
		if err != nil {
			panic(err.Error())
		}

		p.Name = name
		p.Desc = description
		p.Price = price
		p.Quantity = quantity

		products = append(products, p)
	}

	temp.ExecuteTemplate(w, "Index", products)
}
