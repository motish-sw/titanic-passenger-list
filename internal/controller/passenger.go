package controller

import (
	"net/http"

	"github.com/motish-sw/titanic-passenger-list/internal/context"
	"github.com/motish-sw/titanic-passenger-list/internal/models"

	"github.com/labstack/echo/v4"
)

type (
	Passenger          struct{}
	PassengerViewModel struct {
		Name string
		ID   string
	}
)

func (ctrl Passenger) GetPassengerJSON(c echo.Context) error {
	cc := c.(*context.AppContext)
	passengerID := c.Param("id")

	passenger := models.Passenger{ID: passengerID}

	err := cc.PassengerStore.First(&passenger)

	if err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusNotFound, "Passenger ID not found")
	}

	vm := PassengerViewModel{
		Name: passenger.Name,
		ID:   passenger.ID,
	}

	return c.JSON(http.StatusOK, vm)
}
