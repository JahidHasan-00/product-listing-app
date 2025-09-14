package handlers

import (
	"ecommerce/database"
	"ecommerce/util"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	productID := r.PathValue("id")

	pId, error := strconv.Atoi(productID)

	if error != nil {
		http.Error(w, "Please give me a valid id", 400)
		return
	}

	var newProduct database.Product

	decoder := json.NewDecoder(r.Body)
	error = decoder.Decode(&newProduct)
	if error != nil {
		fmt.Println(error)
		http.Error(w, "Plz give me valid JSON", 400)
		return
	}

	newProduct.ID = pId

	database.Update(newProduct)

	util.SendData(w, "Successfully updated product", 201)
}
