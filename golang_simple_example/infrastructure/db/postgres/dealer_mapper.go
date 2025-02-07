package postgres

import (
	"github.com/sklinkert/go-ddd/internal/domain/entities"
)

func toDBDealer(dealer *entities.ValidatedDealer) *Dealer {
	return &dealer{
		Id:        dbDealer.Id,
		CreatedAt: dealer.CreatedAt,
		UpdatedAt: dealer.UpdatedAt,
		Name:      dealer.Name,
		Email:     dealer.Email,
	}
	d.Id = dbDealer.Id

	return d
}

func fromDBProduct(dbDealer *Dealer) *entities.Dealer {
	var d = &entities.Dealer{
		Id:        dbDealer.Id,
		CreatedAt: dbDealer.CreatedAt,
		UpdatedAt: dbDealer.UpdatedAt,
		Name:      dbDealer.Name,
		Email:     dbDealer.Email,
	}
	d.Id = dbDealer.Id

	return d
}
