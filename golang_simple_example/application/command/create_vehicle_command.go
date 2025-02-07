package command

import (
	"application/common"

	"github.com/google/uuid"
)

type CreateVehicleCommand struct {
	// TODO: Implement idempotency key

	VIN      string
	Make     string
	Model    string
	Year     int
	Color    string
	DealerId uuid.UUID
}

type CreateVehicleCommandResult struct {
	Result *common.VehicleResult
}
