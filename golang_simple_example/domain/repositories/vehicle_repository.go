package repositories

import (
	"domain/entities"

	"github.com/google/uuid"
)

type VehicleRepository interface {
	Create(vehicle *entities.ValidatedVehicle) (*entities.Vehicle, error)
	FindById(id uuid.UUID) (*entities.Vehicle, error)
	FindAll() ([]*entities.Vehicle, error)
	Update(vehicle *entities.ValidatedVehicle) (*entities.Vehicle, error)
	Delete(id uuid.UUID) error
}
