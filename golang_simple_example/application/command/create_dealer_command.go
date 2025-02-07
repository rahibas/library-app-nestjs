package command

import "github.com/sklinkert/go-ddd/internal/application/common"

type CreateDealerCommand struct {
	// TODO: Implement idempotency key

	Name string
}

type CreateDealerCommandResult struct {
	Result *common.DealerResult
}
