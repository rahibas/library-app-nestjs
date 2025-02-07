package command

import (
	"github.com/google/uuid"
	"github.com/sklinkert/go-ddd/internal/application/common"
)

type UpdateDealerCommand struct {
	// TODO: Implement idempotency key

	Id   uuid.UUID
	Name string
}

type UpdateDealerCommandResult struct {
	Result *common.DealerResult
}
