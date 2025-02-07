// Client requests the creation of an vehicle-inspection value-object
// HandleCommand is the facade handling Domain Commands, that will eventually trigger registered Event handlers.package factories
// The factory creates the value-object and returns it to the client

package factories

import "time"

func (f *InspectionCreationFactory) CreateInspection(inspection *domain.Inspection) (*domain.Inspection, error) {
	inspection.Id = uuid.New()
	inspection.CreatedAt = time.Now()
	inspection.UpdatedAt = time.Now()

	return inspection, nil
}
