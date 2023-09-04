package context

import (
	"github.com/motish-sw/titanic-passenger-list/config"

	"github.com/motish-sw/titanic-passenger-list/internal/store"

	echo "github.com/labstack/echo/v4"
)

type AppContext struct {
	echo.Context
	PassengerStore store.Passenger
	Config         *config.Configuration
}
