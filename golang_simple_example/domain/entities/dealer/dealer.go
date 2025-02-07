package entities

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Dealer struct {
	Id        string
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      string
	Email     string
}

func (d *Dealer) validate() error {
	if d.Name == "" {
		return errors.New("Dealer name is required")
	}
	if d.Email == "" {
		return errors.New("Dealer email is required")
	}
	return nil
}
func NewDealer(name, email string) *Dealer {
	return &Dealer{
		Id:        uuid.NewString(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      name,
		Email:     email,
	}
}

func (d *Dealer) UpdateDealerName(name string) error {
	d.Name = name
	d.UpdatedAt = time.Now()

	return d.validate()
}

func (d *Dealer) UpdateDealerEmail(email string) error {
	d.Email = email
	d.UpdatedAt = time.Now()

	return d.validate()
}
