package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	Version       string
	ServiceName   string
	HttpPort      int
	JwtSecreteKey string
}

var configuration Config

func loadConfig() {

	error := godotenv.Load()
	if error != nil {
		fmt.Println("Failed to load the env variables")
		os.Exit(1)
	}

	version := os.Getenv("VERSION")
	if version == "" {
		fmt.Println("Version is Required")
		os.Exit(1)
	}

	serviceName := os.Getenv("SERVICE_NAME")
	if serviceName == "" {
		fmt.Println("Service Name is Required")
		os.Exit(1)
	}

	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		fmt.Println("HttpPort is Required")
		os.Exit(1)
	}

	port, error := strconv.ParseInt(httpPort, 10, 64)
	if error != nil {
		fmt.Println("Enter a Valid Port")
		os.Exit(1)
	}

	jwtSecreteKey := os.Getenv("JWT_SECRET_KEY")
	if jwtSecreteKey == "" {
		fmt.Println("SecreteKey is Required")
		os.Exit(1)
	}

	configuration = Config{
		Version:       version,
		ServiceName:   serviceName,
		HttpPort:      int(port),
		JwtSecreteKey: jwtSecreteKey,
	}

}

func GetConfig() Config {
	loadConfig()
	return configuration
}
