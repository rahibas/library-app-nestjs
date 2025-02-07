package services

import (
	"errors"
	"fmt"
	"testing"

	"application/command"
	"domain/entities"

	"github.com/google/uuid"
)

// MockVehicleRepository is a mock implementation of the VehicleRepository interface
type MockVehicleRepository struct {
	vehicles []*entities.ValidatedVehicle
}

func (m *MockVehicleRepository) Create(vehicle *entities.ValidatedVehicle) (*entities.Vehicle, error) {
	m.vehicles = append(m.vehicles, vehicle)
	return &vehicle.Vehicle, nil
}

func (m *MockVehicleRepository) FindAll() ([]*entities.Vehicle, error) {
	var vehicles []*entities.Vehicle
	for _, v := range m.vehicles {
		vehicles = append(vehicles, &v.Vehicle)
	}
	return vehicles, nil
}

func (m *MockVehicleRepository) Update(vehicle *entities.ValidatedVehicle) (*entities.Vehicle, error) {
	for index, p := range m.vehicles {
		if p.Id == vehicle.Id {
			m.vehicles[index] = vehicle
			return &vehicle.Vehicle, nil
		}
	}
	return nil, errors.New("vehicle not found for update")
}

func (m *MockVehicleRepository) Delete(id uuid.UUID) error {
	for index, p := range m.vehicles {
		if p.Id == id {
			m.vehicles = append(m.vehicles[:index], m.vehicles[index+1:]...)
			return nil
		}
	}
	return errors.New("vehicle not found for delete")
}

func (m *MockVehicleRepository) FindById(id uuid.UUID) (*entities.Vehicle, error) {
	for _, p := range m.vehicles {
		if p.Id == id {
			return &p.Vehicle, nil
		}
		fmt.Printf("Id: mem:%s - %s\n", p.Id, id)
	}
	return nil, errors.New("vehicle not found")
}

func TestVehicleService_CreateVehicle(t *testing.T) {
	vehicleRepo := &MockVehicleRepository{}
	dealerRepo := &MockDealerRepository{}
	service := NewVehicleService(vehicleRepo, dealerRepo)

	// Create dealer
	dealer := createPersistedDealer(t, dealerRepo)

	// Create vehicle
	vehicle := entities.NewVehicle("Example", 100.0, *dealer)
	vehicleCommand := getCreateVehicleCommand(vehicle)
	_, err := service.CreateVehicle(vehicleCommand)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}

	if len(vehicleRepo.vehicles) != 1 {
		t.Errorf("Expected 1 vehicle in vehicleRepository, but got %d", len(vehicleRepo.vehicles))
	}
}

func TestVehicleService_GetAllVehicles(t *testing.T) {
	vehicleRepo := &MockVehicleRepository{}
	dealerRepo := &MockDealerRepository{}
	service := NewVehicleService(vehicleRepo, dealerRepo)

	// Create dealer
	dealer := createPersistedDealer(t, dealerRepo)

	// Add two vehicles
	_, _ = service.CreateVehicle(getCreateVehicleCommand(entities.NewVehicle("Example1", 100.0, *dealer)))
	_, _ = service.CreateVehicle(getCreateVehicleCommand(entities.NewVehicle("Example2", 200.0, *dealer)))

	vehicles, err := service.FindAllVehicles()
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}

	if len(vehicles.Result) != 2 {
		t.Errorf("Expected 2 vehicles, but got %d", len(vehicles.Result))
	}
}

func TestVehicleService_FindVehicleById(t *testing.T) {
	vehicleRepo := &MockVehicleRepository{}
	dealerRepo := &MockDealerRepository{}
	service := NewVehicleService(vehicleRepo, dealerRepo)

	// Create dealer
	dealer := createPersistedDealer(t, dealerRepo)

	vehicle := entities.NewVehicle("Example", 100.0, *dealer)
	result, err := service.CreateVehicle(getCreateVehicleCommand(vehicle))
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}

	foundVehicle, err := service.FindVehicleById(result.Result.Id)
	if err != nil {
		t.Fatalf("Unexpected error: %s", err)
	}

	if foundVehicle.Result.Name != "Example" {
		t.Errorf("Expected vehicle name 'Example', but got %s", foundVehicle.Result.Name)
	}

	_, err = service.FindVehicleById(uuid.New()) // some non-existent Id
	if err == nil {
		t.Error("Expected error for non-existent vehicle, but got none")
	}
}

func getCreateVehicleCommand(vehicle *entities.Vehicle) *command.CreateVehicleCommand {
	return &command.CreateVehicleCommand{
		Name:     vehicle.Name,
		Price:    vehicle.Price,
		DealerId: vehicle.Dealer.Id,
	}
}

func createPersistedDealer(t *testing.T, dealerRepo *MockDealerRepository) *entities.ValidatedDealer {
	dealer := entities.NewDealer("John Doe")
	validatedDealer, err := entities.NewValidatedDealer(dealer)
	if err != nil {
		t.Fatalf("Unexpected error: %s", err)
	}
	_, err = dealerRepo.Create(validatedDealer)
	if err != nil {
		t.Fatalf("Unexpected error: %s", err)
	}
	return validatedDealer
}
