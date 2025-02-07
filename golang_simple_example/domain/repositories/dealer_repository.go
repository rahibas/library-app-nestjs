package repositories

import (
	"domain/entities"

	"github.com/google/uuid"
)

type DealerRepository interface {
	Create(dealer *entities.ValidatedDealer) (*entities.Dealer, error)
	FindById(id uuid.UUID) (*entities.Dealer, error)
	FindAll() ([]*entities.Dealer, error)
	Update(dealer *entities.ValidatedDealer) (*entities.Dealer, error)
	Delete(id uuid.UUID) error
}
