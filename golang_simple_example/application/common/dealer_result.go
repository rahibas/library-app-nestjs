package common

import (
	"time"

	"github.com/google/uuid"
)

type DealerResult struct {
	Id        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      string
	Email     string
}
