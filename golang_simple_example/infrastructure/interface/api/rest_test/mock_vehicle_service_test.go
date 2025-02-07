package rest_test

import (
	"time"

	"github.com/google/uuid"
	"github.com/sklinkert/go-ddd/internal/application/command"
	"github.com/sklinkert/go-ddd/internal/application/mapper"
	"github.com/sklinkert/go-ddd/internal/application/query"
	"github.com/sklinkert/go-ddd/internal/domain/entities"
	"github.com/stretchr/testify/mock"
)

type MockVehicleService struct {
	mock.Mock
}

func (m *MockVehicleService) CreateVehicle(vehicleCommand *command.CreateVehicleCommand) (*command.CreateVehicleCommandResult, error) {
	args := m.Called(vehicleCommand)

	var now = time.Now()

	var dealer = &entities.Dealer{
		Id:        vehicleCommand.DealerId,
		Name:      "Test Dealer",
		CreatedAt: now,
		UpdatedAt: now,
	}

	var validatedDealer, err = entities.NewValidatedDealer(dealer)
	if err != nil {
		return nil, err
	}

	var newVehicle = entities.NewVehicle(
		vehicleCommand.Name,
		vehicleCommand.Price,
		*validatedDealer,
	)

	validatedVehicle, err := entities.NewValidatedVehicle(newVehicle)
	if err != nil {
		return nil, err
	}

	var result command.CreateVehicleCommandResult
	result.Result = mapper.NewVehicleResultFromValidatedEntity(validatedVehicle)

	return &result, args.Error(1)
}

func (m *MockVehicleService) FindAllVehicles() (*query.VehicleQueryListResult, error) {
	args := m.Called()

	vehicleQueryListResult := &query.VehicleQueryListResult{}

	for _, vehicle := range args.Get(0).([]*entities.Vehicle) {
		vehicleQueryListResult.Result = append(vehicleQueryListResult.Result, mapper.NewVehicleResultFromEntity(vehicle))
	}

	return vehicleQueryListResult, args.Error(1)
}

func (m *MockVehicleService) FindVehicleById(id uuid.UUID) (*query.VehicleQueryResult, error) {
	args := m.Called(id)

	vehicleQueryResult := &query.VehicleQueryResult{
		Result: mapper.NewVehicleResultFromEntity(args.Get(0).(*entities.Vehicle)),
	}

	return vehicleQueryResult, args.Error(1)
}
