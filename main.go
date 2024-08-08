package main

import (
	"errors"
	"log"
	"os"

	_ "github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"

	"github.com/joho/godotenv"

	"chatProject/routes/index"
)

type Config struct {
	port string
}

func configureRoutes(e *echo.Echo) {
	index.ConfigureRoutes(e)
}

func getConfig() (*Config, error) {
	port, portSet := os.LookupEnv("APP_PORT")
	if !portSet {
		return nil, errors.New("APP_PORT not found")
	}

	return &Config{
		port: ":" + port,
	}, nil
}

func main() {
	godotenv.Load()

	config, err := getConfig()

	if err != nil {
		log.Fatal(err)
	}

	e := echo.New()

	e.Logger.Fatal(e.Start(config.port))
}
