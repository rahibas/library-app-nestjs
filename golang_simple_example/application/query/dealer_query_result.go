package query

import (
	"application/common"
)

type DealerQueryResult struct {
	Result *common.DealerResult
}

type DealerQueryListResult struct {
	Result []*common.DealerResult
}
