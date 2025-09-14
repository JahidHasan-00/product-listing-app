package handlers

import (
	"ecommerce/database"
	"ecommerce/util"
	"encoding/json"
	"fmt"
	"net/http"
)

type ReqLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Login(w http.ResponseWriter, r *http.Request) {

	var reqLogin ReqLogin

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&reqLogin)

	if err != nil {
		fmt.Println(err)
		http.Error(w, "Invalid Request Data", http.StatusBadRequest)
	}

	user := database.Find(reqLogin.Email, reqLogin.Password)

	if user == nil {
		http.Error(w, "Invalid Credentials", http.StatusBadRequest)
		return
	}

	util.SendData(w, user, http.StatusCreated)
}
