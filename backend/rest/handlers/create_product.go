package handlers

import (
	"ecommerce/database"
	"ecommerce/util"
	"encoding/json"
	"fmt"
	"net/http"
)

func CreateProduct(w http.ResponseWriter, r *http.Request) {

	var newProduct database.Product

	decoder := json.NewDecoder(r.Body)
	error := decoder.Decode(&newProduct)
	if error != nil {
		fmt.Println(error)
		http.Error(w, "Plz give me valid JSON", 400)
		return
	}

	// set ID for new product
	createProduct := database.Store(newProduct)

	//new product to the productLis
	util.SendData(w, createProduct, 201)
}
