package repositories

import (
	"github.com/google/uuid"
	"github.com/sklinkert/go-ddd/internal/domain/entities"
)

type VehicleRepository interface {
	Create(vehicle *entities.ValidatedVehicle) (*entities.Vehicle, error)
	FindById(id uuid.UUID) (*entities.Vehicle, error)
	FindAll() ([]*entities.Vehicle, error)
	Update(vehicle *entities.ValidatedVehicle) (*entities.Vehicle, error)
	Delete(id uuid.UUID) error
}
