package cmd

import (
	"ecommerce/config"
	"ecommerce/rest"
)

func Serve() {

	cnf := config.GetConfig()
	rest.Server(cnf)
}
