package controller

import (
	"net/http"

	"github.com/motish-sw/titanic-passenger-list/internal/context"
	"github.com/motish-sw/titanic-passenger-list/internal/models"

	"github.com/labstack/echo/v4"
)

type (
	PassengerList          struct{}
	PassengerListViewModel struct {
		Passengers []PassengerViewModel
	}
)

func (ctrl PassengerList) GetPassengers(c echo.Context) error {
	cc := c.(*context.AppContext)

	passengers := []models.Passenger{}

	err := cc.PassengerStore.Find(&passengers)

	if err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusNotFound, "Passenger list couldn't be found.")
	}

	viewModel := PassengerListViewModel{
		Passengers: make([]PassengerViewModel, len(passengers)),
	}

	for index, Passenger := range passengers {
		viewModel.Passengers[index] = PassengerViewModel{
			Name: Passenger.Name,
			ID:   Passenger.ID,
		}
	}

	return c.JSON(http.StatusOK, viewModel)

}
