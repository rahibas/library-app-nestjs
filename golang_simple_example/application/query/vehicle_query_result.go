package query

import "application/common"

type VehicleQueryResult struct {
	Result *common.VehicleResult
}

type VehicleQueryListResult struct {
	Result []*common.VehicleResult
}
