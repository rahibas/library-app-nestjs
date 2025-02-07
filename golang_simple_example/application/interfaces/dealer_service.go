package interfaces

import (
	"github.com/google/uuid"
	"github.com/sklinkert/go-ddd/internal/application/command"
	"github.com/sklinkert/go-ddd/internal/application/query"
)

type DealerService interface {
	CreateDealer(dealerCommand *command.CreateDealerCommand) (*command.CreateDealerCommandResult, error)
	FindAllDealers() (*query.DealerQueryListResult, error)
	FindDealerById(id uuid.UUID) (*query.DealerQueryResult, error)
	UpdateDealer(updateCommand *command.UpdateDealerCommand) (*command.UpdateDealerCommandResult, error)
	DeleteDealer(id uuid.UUID) error
}
