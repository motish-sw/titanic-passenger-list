package core

import (
	"titanic-passangers/internal/context"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	mid "titanic-passangers/internal/core/middleware"

	v "gopkg.in/go-playground/validator.v9"
)

func NewRouter(server *Server) *echo.Echo {
	config := server.config
	e := echo.New()
	e.Validator = &Validator{validator: v.New()}

	cc := context.AppContext{
		PassengerStore: &PassengerStore{DB: server.db},
		Config:         config,
	}

	e.Use(mid.AppContext(&cc))

	e.Use(middleware.Recover())       // panic errors are thrown
	e.Use(middleware.BodyLimit("5M")) // limit body payload to 5MB
	e.Use(middleware.Secure())        // provide protection against injection attacks
	e.Use(middleware.RequestID())     // generate unique requestId

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET},
	}))

	return e
}
