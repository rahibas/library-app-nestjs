package interfaces

import (
	"github.com/google/uuid"
	"github.com/sklinkert/go-ddd/internal/application/command"
	"github.com/sklinkert/go-ddd/internal/application/query"
)

type VehicleService interface {
	CreateVehicle(vehicleCommand *command.CreateVehicleCommand) (*command.CreateVehicleCommandResult, error)
	FindAllVehicles() (*query.VehicleQueryListResult, error)
	FindVehicleById(id uuid.UUID) (*query.VehicleQueryResult, error)
}
