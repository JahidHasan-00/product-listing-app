package middlewares

import "net/http"

func Cors(next http.Handler) http.Handler {

	//Handle Cors
	handleAllReq := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, habib")
		w.Header().Set("Content-Type", "application/json")

		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(handleAllReq)
}
