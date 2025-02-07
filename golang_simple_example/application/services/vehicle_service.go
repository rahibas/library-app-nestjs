package services

import (
	"errors"

	"application/command"    // Replace this with the correct import path
	"application/interfaces" // Replace this with the correct import path
	"application/mapper"     // Replace this with the correct import path
	"application/query"      // Replace this with the correct import path
	"domain/entities"        // Replace this with the correct import path
	"domain/repositories"    // Replace this with the correct import path

	"github.com/google/uuid"
)

type VehicleService struct {
	vehicleRepository repositories.VehicleRepository
	dealerRepository  repositories.DealerRepository
}

// NewVehicleService - Constructor for the service
func NewVehicleService(
	vehicleRepository repositories.VehicleRepository,
	dealerRepository repositories.DealerRepository,
) interfaces.VehicleService {
	return &VehicleService{vehicleRepository: vehicleRepository, dealerRepository: dealerRepository}
}

// CreateVehicle saves a new vehicle
func (v *VehicleService) CreateVehicle(vehicleCommand *command.CreateVehicleCommand) (*command.CreateVehicleCommandResult, error) {
	storedDealer, err := v.dealerRepository.FindById(vehicleCommand.DealerId)
	if err != nil {
		return nil, err
	}

	if storedDealer == nil {
		return nil, errors.New("Dealer not found")
	}

	validatedDealer, err := entities.NewValidatedDealer(*storedDealer)
	if err != nil {
		return nil, err
	}

	var newVehicle = entities.NewVehicle(
		vehicleCommand.Make,
		vehicleCommand.Model,
		vehicleCommand.Year,
		vehicleCommand.Color,
		vehicleCommand.DealerId,
		*validatedDealer,
	)

	validatedVehicle, err := entities.NewValidatedVehicle(newVehicle)
	if err != nil {
		return nil, err
	}

	_, err = v.vehicleRepository.Create(validatedVehicle)
	if err != nil {
		return nil, err
	}

	result := command.CreateVehicleCommandResult{
		Result: mapper.NewVehicleResultFromValidatedEntity(validatedVehicle),
	}

	return &result, nil
}

// FindAllVehicles fetches all vehicles
func (v *VehicleService) FindAllVehicless() (*query.VehicleQueryListResult, error) {
	storedVehicles, err := v.vehicleRepository.FindAll()
	if err != nil {
		return nil, err
	}

	var queryListResult query.VehicleQueryListResult
	for _, vehicle := range storedVehicles {
		queryListResult.Result = append(queryListResult.Result, mapper.NewProductResultFromEntity(vehicle))
	}

	return &queryListResult, nil
}

// FindVehicleById fetches a specific vehicle by Id
func (v *VehicleService) FindVehicleById(id uuid.UUID) (*query.VehicleQueryResult, error) {
	storedVehicle, err := v.vehicleRepository.FindById(id)
	if err != nil {
		return nil, err
	}

	var queryResult query.VehicleQueryResult
	queryResult.Result = mapper.NewVehiclesResultFromEntity(storedVehicle)

	return &queryResult, nil
}
