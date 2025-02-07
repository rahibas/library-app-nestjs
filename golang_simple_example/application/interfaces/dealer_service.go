package interfaces

import (
	"application/command"
	"application/query"

	"github.com/google/uuid"
)

type DealerService interface {
	CreateDealer(dealerCommand *command.CreateDealerCommand) (*command.CreateDealerCommandResult, error)
	FindAllDealers() (*query.DealerQueryListResult, error)
	FindDealerById(id uuid.UUID) (*query.DealerQueryResult, error)
	UpdateDealer(updateCommand *command.UpdateDealerCommand) (*command.UpdateDealerCommandResult, error)
	DeleteDealer(id uuid.UUID) error
}
