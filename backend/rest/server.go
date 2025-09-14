package rest

import (
	"ecommerce/config"
	middleware "ecommerce/rest/middlewares"
	"fmt"
	"net/http"
	"strconv"
)

func Server(cnf config.Config) {

	manager := middleware.NewManager()

	mux := http.NewServeMux() // mux is a router
	manager.Use(
		middleware.Logger,
		middleware.Preflight,
		middleware.Cors,
	)

	wrappedMux := manager.WrapMux(mux)

	initRoutes(mux, manager)

	add := ":" + strconv.Itoa(cnf.HttpPort)

	fmt.Println("Server running on", add)
	err := http.ListenAndServe(add, wrappedMux) // "Failed to start the server"

	if err != nil {
		fmt.Println("Error starting the server", err)
	}
}
