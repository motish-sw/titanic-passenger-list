package store

import "github.com/motish-sw/titanic-passenger-list/internal/models"

type Passenger interface {
	First(m *models.Passenger) error
	Find(m *[]models.Passenger) error
}
