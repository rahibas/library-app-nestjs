package command

import (
	"application/common"

	"github.com/google/uuid"
)

type UpdateDealerCommand struct {
	// TODO: Implement idempotency key

	Id    uuid.UUID
	Name  string
	Email string
}

type UpdateDealerCommandResult struct {
	Result *common.DealerResult
}
