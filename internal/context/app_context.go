package context

import (
	"titanic-passengers/config"

	"titanic-passengers/internal/store"

	echo "github.com/labstack/echo/v4"
)

type AppContext struct {
	echo.Context
	PassengerStore store.Passenger
	Config         *config.Configuration
}
