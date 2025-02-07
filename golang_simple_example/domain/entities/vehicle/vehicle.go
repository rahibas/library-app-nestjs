package entities

import (
	"errors"
	"time"
)

type Vehicle struct {
	VIN       string
	CreatedAt time.Time
	UpdatedAt time.Time
	Make      string
	Model     string
	Year      int
	Color     string
}

func (v *Vehicle) validate() error {
	if v.VIN == "" {
		return errors.New("Vehicle VIN is required")
	}
	if v.Make == "" {
		return errors.New("Vehicle make is required")
	}
	if v.Model == "" {
		return errors.New("Vehicle model is required")
	}
	if v.Year == 0 {
		return errors.New("Vehicle year is required")
	}
	if v.Color == "" {
		return errors.New("Vehicle color is required")
	}
	return nil
}
