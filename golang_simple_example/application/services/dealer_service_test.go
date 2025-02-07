package services

import (
	"errors"
	"fmt"
	"testing"

	"application/command"
	"domain/entities"

	"github.com/google/uuid"
)

// MockDealerRepository is a mock implementation of the DealerRepository interface
type MockDealerRepository struct {
	dealers []*entities.ValidatedDealer
}

func (m *MockDealerRepository) Create(dealer *entities.ValidatedDealer) (*entities.Dealer, error) {
	m.dealers = append(m.dealers, dealer)
	return &dealer.Dealer, nil
}

func (m *MockDealerRepository) FindAll() ([]*entities.Dealer, error) {
	var dealers []*entities.Dealer
	for _, s := range m.dealers {
		dealers = append(dealers, &s.Dealer)
	}
	return dealers, nil
}

func (m *MockDealerRepository) FindById(id uuid.UUID) (*entities.Dealer, error) {
	for _, s := range m.dealers {
		if s.Id == id {
			return &s.Dealer, nil
		} else {
			fmt.Printf("Id: %s - %s\n", s.Id, id)
		}
	}
	return nil, errors.New("dealer not found")
}

func (m *MockDealerRepository) Delete(id uuid.UUID) error {
	for index, s := range m.dealers {
		if s.Id == id {
			m.dealers = append(m.dealers[:index], m.dealers[index+1:]...)
			return nil
		}
	}
	return errors.New("dealer not found for deletion")
}

func (m *MockDealerRepository) Update(dealer *entities.ValidatedDealer) (*entities.Dealer, error) {
	for index, s := range m.dealers {
		if s.Id == dealer.Id {
			m.dealers[index] = dealer
			return &dealer.Dealer, nil
		}
	}
	return nil, errors.New("dealer not found for update")
}

func TestDealerService_CreateDealer(t *testing.T) {
	repo := &MockDealerRepository{}
	service := NewDealerService(repo)

	_, err := service.CreateDealer(getCreateDealerCommand("John Doe"))
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}

	if len(repo.dealers) != 1 {
		t.Errorf("Expected 1 dealer in vehicleRepository, but got %d", len(repo.dealers))
	}
}

func TestDealerService_GetAllDealers(t *testing.T) {
	repo := &MockDealerRepository{}
	service := NewDealerService(repo)

	// Add two dealers
	_, _ = service.CreateDealer(getCreateDealerCommand("John Doe"))
	_, _ = service.CreateDealer(getCreateDealerCommand("Jane Doe"))

	dealers, err := service.FindAllDealers()
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}

	if len(dealers.Result) != 2 {
		t.Errorf("Expected 2 dealers, but got %d", len(dealers.Result))
	}
}

func TestDealerService_GetDealerById(t *testing.T) {
	repo := &MockDealerRepository{}
	service := NewDealerService(repo)

	createdDealerResult, _ := service.CreateDealer(getCreateDealerCommand("John Doe"))
	dealerID := createdDealerResult.Result.Id

	foundDealer, err := service.FindDealerById(dealerID)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}

	if foundDealer.Result.Name != "John Doe" {
		t.Errorf("Expected dealer name 'John Doe', but got %s", foundDealer.Result.Name)
	}

	_, err = service.FindDealerById(uuid.New()) // some non-existent Id
	if err == nil {
		t.Error("Expected error for non-existent dealer, but got none")
	}
}

func TestDealerService_UpdateDealer(t *testing.T) {
	repo := &MockDealerRepository{}
	service := NewDealerService(repo)

	createdDealerResult, _ := service.CreateDealer(getCreateDealerCommand("John Doe"))
	dealerId := createdDealerResult.Result.Id

	var updatableDealer = entities.Dealer{
		Id:   dealerId,
		Name: "Doe Johnny",
	}

	_, err := service.UpdateDealer(&command.UpdateDealerCommand{
		Id:   dealerId,
		Name: updatableDealer.Name,
	})
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}

	updatedDealer, _ := service.FindDealerById(dealerId)
	if updatedDealer.Result.Name != "Doe Johnny" {
		t.Errorf("Expected dealer name 'Johnny Doe', but got %s", updatedDealer.Result.Name)
	}
}

func getCreateDealerCommand(name string) *command.CreateDealerCommand {
	return &command.CreateDealerCommand{
		Name: name,
	}
}
