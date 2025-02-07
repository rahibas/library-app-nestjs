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

type DealerService struct {
	repo repositories.DealerRepository
}

// NewDealerService - Constructor for the service
func NewDealerService(repo repositories.DealerRepository) interfaces.DealerService {
	return &DealerService{repo: repo}
}

// CreateDealer saves a new dealer
func (d *DealerService) CreateDealer(dealerCommand *command.CreateDealerCommand) (*command.CreateDealerCommandResult, error) {
	var newDealer = entities.NewDealer(dealerCommand.Name)

	validatedDealer, err := entities.NewValidatedDealer(newDealer)
	if err != nil {
		return nil, err
	}

	_, err = d.repo.Create(validatedDealer)
	if err != nil {
		return nil, err
	}

	result := command.CreateDealerCommandResult{
		Result: mapper.NewDealerResultFromValidatedEntity(validatedDealer),
	}

	return &result, nil
}

// FindAllDealers fetches all sellers
func (d *DealerService) FindAllDealers() (*query.DealerQueryListResult, error) {
	storedDealers, err := d.repo.FindAll()
	if err != nil {
		return nil, err
	}

	var queryResult query.DealerQueryListResult
	for _, dealer := range storedDealers {
		queryResult.Result = append(queryResult.Result, mapper.NewDealerResultFromEntity(dealer))
	}

	return &queryResult, nil
}

// FindDealerById fetches a specific dealer by Id
func (d *DealerService) FindDealerById(id uuid.UUID) (*query.DealerQueryResult, error) {
	storedDealer, err := d.repo.FindById(id)
	if err != nil {
		return nil, err
	}

	var queryResult query.DealerQueryResult
	queryResult.Result = mapper.NewDealerResultFromEntity(storedDealer)

	return &queryResult, nil
}

// UpdateDealer updates a dealer
func (d *DealerService) UpdateDealer(updateCommand *command.UpdateDealerCommand) (*command.UpdateDealerCommandResult, error) {
	dealer, err := d.repo.FindById(updateCommand.Id)
	if err != nil {
		return nil, err
	}

	if dealer == nil {
		return nil, errors.New("dealer not found")
	}

	if err := dealer.UpdateName(updateCommand.Name); err != nil {
		return nil, err
	}

	validatedUpdatedDealer, err := entities.NewValidatedDealer(dealer)
	if err != nil {
		return nil, err
	}

	_, err = d.repo.Update(validatedUpdatedDealer)
	if err != nil {
		return nil, err
	}

	result := command.UpdateDealerCommandResult{
		Result: mapper.NewDealerResultFromEntity(dealer),
	}

	return &result, nil
}

// DeleteDealer deletes a dealer
func (d *DealerService) DeleteDealer(id uuid.UUID) error {
	return d.repo.Delete(id)
}
