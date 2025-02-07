// Client requests the creation of an vehicle-inspection value-object
// HandleCommand is the facade handling Domain Commands, that will eventually trigger registered Event handlers.package factories
// The factory creates the value-object and returns it to the client

package factories

import "time"

func (f *DealerInventoryCreationFactory) CreateDealerInventory(dealerInventory *domain.DealerInventory) (*domain.DealerInventory, error) {
	dealerInventory.Id = uuid.New()
	dealerInventory.CreatedAt = time.Now()
	dealerInventory.UpdatedAt = time.Now()

	return dealerInventory, nil
}
