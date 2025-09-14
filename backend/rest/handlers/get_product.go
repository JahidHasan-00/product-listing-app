package handlers

import (
	"ecommerce/database"
	"ecommerce/util"
	"net/http"
	"strconv"
)

func GetProduct(w http.ResponseWriter, r *http.Request) {
	productID := r.PathValue("id")
	pId, error := strconv.Atoi(productID)

	if error != nil {
		http.Error(w, "Please give me a valid product id", 400)
		return
	}

	product := database.Get(pId)

	if product == nil {
		util.SendError(w, 404, "Product Not Found")
		return
	}

	util.SendData(w, product, 200)

}
