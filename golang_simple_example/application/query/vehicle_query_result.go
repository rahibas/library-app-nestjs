package query

import "github.com/sklinkert/go-ddd/internal/application/common"

type VehicleQueryResult struct {
	Result *common.VehicleResult
}

type VehicleQueryListResult struct {
	Result []*common.VehicleResult
}
