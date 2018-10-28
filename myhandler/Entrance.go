package myhandler

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

func Entrance(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("layout.html", "tax_calculation.html")
	if err != nil {
		fmt.Println("wtf")
		fmt.Println(err)
	}
	data := map[string]float64{}
	price, err := strconv.ParseFloat(r.FormValue("price"), 64)
	if err != nil {
		price = 0.0
	}
	isWithTax, err := strconv.ParseBool(r.FormValue("is_with_tax"))
	if err != nil {
		isWithTax = false
	}

	fmt.Println("price:", price, "\nisWithTax:", isWithTax)
	var taxParcentage = 0.08
	if isWithTax {
		data["with_tax_price"] = price
		data["no_tax_price"] = price / (1 + taxParcentage)
		data["tax"] = taxParcentage * 100
		data["tax_price"] = taxParcentage * data["no_tax_price"]
	} else {
		data["with_tax_price"] = price * (1 + taxParcentage)
		data["no_tax_price"] = price
		data["tax"] = taxParcentage * 100
		data["tax_price"] = taxParcentage * data["no_tax_price"]
	}
	t.ExecuteTemplate(w, "layout", data)
}
