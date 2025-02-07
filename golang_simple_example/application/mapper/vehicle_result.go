package mapper

import (
	"application/common"
	"domain/entities"
)

func NewVehicleResultFromValidatedEntity(vehicle *entities.ValidatedVehicle) *common.VehicleResult {
	return NewVehicleResultFromEntity(&vehicle.Vehicle)
}

func NewVehicleResultFromEntity(vehicle *entities.Vehicle) *common.VehicleResult {
	if vehicle == nil {
		return nil
	}

	return &common.VehicleResult{
		Id:        vehicle.VIN,
		Name:      vehicle.Name,
		Make:      vehicle.Make,
		Model:     vehicle.Model,
		Year:      vehicle.Year,
		Color:     vehicle.Color,
		Dealer:    NewDealerResultFromEntity(&vehicle.Dealer),
		CreatedAt: vehicle.CreatedAt,
		UpdatedAt: vehicle.UpdatedAt,
	}
}
