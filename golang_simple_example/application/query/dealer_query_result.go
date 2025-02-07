package query

import (
	"github.com/sklinkert/go-ddd/internal/application/common"
)

type DealerQueryResult struct {
	Result *common.DealerResult
}

type DealerQueryListResult struct {
	Result []*common.DealerResult
}
