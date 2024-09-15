package auth

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"

	"chatProject/routes/shared"

	"github.com/gorilla/sessions"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type route struct {
	db *pgxpool.Pool
}

type User struct {
	Username        string `form:"username" json:"username"`
	Password        string `form:"password" json:"password"`
	ConfirmPassword string `form:"confirm_password" json:"confirm_password"`
	ID              int32
}

func ConfigureRoutes(e *echo.Echo, db *pgxpool.Pool) {
	routes := route{db: db}

	e.GET("/login", routes.login)
	e.GET("/register", Auth(routes.register))
	e.GET("/test2", Auth(routes.test))

	e.POST("/login", routes.login)
	e.POST("/register", routes.register)

}

func (route *route) test(ctx echo.Context) error {
	log.Println("I was here bro")
	return ctx.String(200, "Salam!")
}

func Auth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		session, err := session.Get("session", ctx)

		if err != nil {
			return err
		}

		user_id, ok := session.Values["user_id"]
		if !ok {
			fmt.Println("Unauthorized")
			return ctx.Redirect(http.StatusTemporaryRedirect, "/login")
		}
		ctx.Set("user_id", user_id)
		return next(ctx)
	}
}

func (route *route) login(ctx echo.Context) error {
	if ctx.Request().Method == "POST" {
		form := new(User)
		if err := ctx.Bind(form); err != nil {
			return err
		}
		if form.Username == "" {
			log.Println("Username is empty")
			return nil
		}
		if form.Password == "" {
			log.Println("Password is empty")
			return nil
		}

		user, err := getUser(route.db, form.Username)

		if err != nil {
			return err
		}
		// # Ensure username exists and password is correct
		if user.Username == form.Username {
			if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(form.Password)) != nil {
				fmt.Println("Pass is not correct")
				return shared.Page(ctx, Login())
			}
			fmt.Println("Logged in")
		}

		sess, err := session.Get("session", ctx)
		if err != nil {
			return err
		}

		sess.Options = &sessions.Options{
			Path:     "/",
			MaxAge:   86400 * 7,
			HttpOnly: true,
		}

		sess.Values["user_id"] = user.ID

		if err := sess.Save(ctx.Request(), ctx.Response()); err != nil {
			return err
		}

		return ctx.Redirect(http.StatusTemporaryRedirect, "/")
	} else {
		return shared.Page(ctx, Login())
	}
}

func (route *route) register(ctx echo.Context) error {
	if ctx.Request().Method == "POST" {
		form := new(User)
		if err := ctx.Bind(form); err != nil {
			return err
		}
		// Validation
		if form.Password == "" {
			log.Println("Pass is empty")
			return nil
		} else if len(strings.TrimSpace(form.Password)) == 0 {
			log.Println("Blank is not supported")
			return nil
		} else if form.Password != form.ConfirmPassword {
			log.Println("Pass don't match")
			return nil
		} else if len(form.Password) < 8 {
			log.Println("Password must be at least 8 characters long")
			return nil
		} else if len(form.Password) > 64 {
			log.Println("Password must be lower than 64 characters long")
			return nil
		}
		if form.Username == "" {
			log.Println("Username is empty")
			return nil
		} else if len(strings.TrimSpace(form.Username)) != len(form.Username) {
			log.Println("Blank is not supported")
			return nil
		}

		err := createUser(route.db, form.Username, form.Password)
		if err != nil {
			return err
		}

		return ctx.Redirect(http.StatusTemporaryRedirect, "/")
	} else {
		ctx.Logger()
		return shared.Page(ctx, Registration())
	}
}

type UserFetchError struct {
	username string
}

func (e *UserFetchError) Error() string {
	return fmt.Sprintf("Couldn't find user by username %v", e.username)
}

func getUser(db *pgxpool.Pool, username string) (*User, error) {
	rows, err := db.Query(context.Background(), "SELECT name, pass, id FROM test WHERE name = @name", pgx.NamedArgs{
		"name": username,
	})

	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			return nil, err
		}
	}

	user := new(User)
	len := 0

	for rows.Next() {
		len++
		values, err := rows.Values()
		if err != nil {
			return nil, err
		}
		user.Username = values[0].(string)
		user.Password = values[1].(string)
		user.ID = values[2].(int32)
	}

	if len == 0 {
		return nil, &UserFetchError{username}
	}
	return user, nil
}

func createUser(db *pgxpool.Pool, username string, password string) error {
	hashPass, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return err
		//TO DO: Refactor
	}

	_, err = db.Exec(context.Background(), "INSERT INTO test (name, pass) VALUES (@name, @pass)", pgx.NamedArgs{
		"name": username,
		"pass": hashPass,
	})

	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			log.Println("Username already exist")
			return err
		}
	}
	return nil

}
