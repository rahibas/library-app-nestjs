package interfaces

import (
	"application/command"
	"application/query"

	"github.com/google/uuid"
)

type VehicleService interface {
	CreateVehicle(vehicleCommand *command.CreateVehicleCommand) (*command.CreateVehicleCommandResult, error)
	FindAllVehicles() (*query.VehicleQueryListResult, error)
	FindVehicleById(id uuid.UUID) (*query.VehicleQueryResult, error)
}
