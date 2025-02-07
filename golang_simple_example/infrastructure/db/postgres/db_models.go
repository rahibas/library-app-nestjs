package postgres

import (
	"time"

	"github.com/google/uuid"
)

type Dealer struct {
	Id        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      string
	Email     string
}

type Inventory struct{}
type Inspection struct{}
type Vehicle struct{}
