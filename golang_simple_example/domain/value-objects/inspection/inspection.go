package value_objects

func (i VehicleInspection) IsValid() bool {
	// Compare this snippet from golang_simple_example/domain/value-objects/inspection/inspection.go:
	// 	return i.Id != "" && i.VehicleId != "" && i.DealerId != "" && i.InspectionDate != "" && i.InspectionType != "" && i.InspectionResult != ""
	return i.Id != "" && i.VehicleId != "" && i.DealerId != "" && i.InspectionDate != "" && i.InspectionType != "" && i.InspectionResult != ""
}
