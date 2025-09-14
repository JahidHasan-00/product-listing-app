package handlers

import (
	"ecommerce/database"
	"ecommerce/util"
	"net/http"
	"strconv"
)

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	productID := r.PathValue("id")

	pId, error := strconv.Atoi(productID)

	if error != nil {
		http.Error(w, "Please give me a valid id", 400)
		return
	}

	database.Delete(pId)

	util.SendData(w, "Successfully deleted product", 201)
}
