package main

import (
	"errors"
	"log"
	"os"

	"github.com/gorilla/sessions"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"

	"chatProject/models"
	"chatProject/routes/auth"
	"chatProject/routes/index"
)

type Config struct {
	port   string
	dbConn string
}

func configureRoutes(e *echo.Echo, db *pgxpool.Pool) {
	e.Static("/public", "public")
	index.ConfigureRoutes(e)
	auth.ConfigureRoutes(e, db)
}

func getConfig() (*Config, error) {
	port, portSet := os.LookupEnv("APP_PORT")
	if !portSet {
		return nil, errors.New("APP_PORT not found")
	}
	dbConn, connSet := os.LookupEnv("DB_CONNECTION_STRING")
	if !connSet {
		return nil, errors.New("DB_CONNECTION_STRING not found")
	}

	return &Config{
		port:   ":" + port,
		dbConn: dbConn,
	}, nil
}

// func seymurgay(next echo.HandlerFunc) echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		if err := next(c); err != nil {
// 			return err
// 		}
// 		session, err := session.Get("session", c)

// 		if err != nil {
// 			return err
// 		}
// 		fmt.Println(session.Values["user_id"])
// 		fmt.Println(c.Request().URL)
// 		return nil
// 	}

// }

func main() {
	godotenv.Load()

	e := echo.New()

	conf, err := getConfig()
	if err != nil {
		log.Fatal(err)
	}

	db, err := models.GetDatabaseConnection(conf.dbConn)
	if err != nil {
		log.Fatal(err)
	}
	secretKey := "sssssqwerty"
	e.Use(session.Middleware(sessions.NewCookieStore([]byte(secretKey))))

	// e.Use(seymurgay)

	configureRoutes(e, db)

	e.Logger.Fatal(e.Start(conf.port))
}
