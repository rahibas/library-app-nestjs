package repositories

import (
	"github.com/google/uuid"
	"github.com/sklinkert/go-ddd/internal/domain/entities"
)

type DealerRepository interface {
	Create(dealer *entities.ValidatedDealer) (*entities.Dealer, error)
	FindById(id uuid.UUID) (*entities.Dealer, error)
	FindAll() ([]*entities.Dealer, error)
	Update(dealer *entities.ValidatedDealer) (*entities.Dealer, error)
	Delete(id uuid.UUID) error
}
