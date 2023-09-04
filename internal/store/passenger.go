package store

import "titanic-passengers/internal/models"

type Passenger interface {
	First(m *models.Passenger) error
	Find(m *[]models.Passenger) error
}
