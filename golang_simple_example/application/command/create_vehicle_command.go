package command

import (
	"github.com/google/uuid"
	"github.com/sklinkert/go-ddd/internal/application/common"
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
