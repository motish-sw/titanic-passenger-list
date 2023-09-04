package core

import (
	"titanic-passangers/internal/models"

	"github.com/jinzhu/gorm"
)

type PassengerStore struct {
	DB *gorm.DB
}

func (s *PassengerStore) First(m *models.Passenger) error {
	return s.DB.First(m).Error
}

func (s *PassengerStore) Find(m *[]models.Passenger) error {
	return s.DB.Find(m).Error
}
