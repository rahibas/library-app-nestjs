package common

import (
	"time"
)

type VehicleResult struct {
	VIN       string
	Make      string
	Model     string
	Year      int
	Color     string
	Dealer    *DealerResult
	CreatedAt time.Time
	UpdatedAt time.Time
}
