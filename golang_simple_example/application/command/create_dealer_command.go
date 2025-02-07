package command

import (
	"application/common"

	"github.com/google/uuid"
)

type CreateDealerCommand struct {
	// TODO: Implement idempotency key
	Id    uuid.UUID
	Name  string
	Email string
}

type CreateDealerCommandResult struct {
	Result *common.DealerResult
}
