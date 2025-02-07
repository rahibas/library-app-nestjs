package mapper

import (
	"github.com/sklinkert/go-ddd/internal/application/common"
	"github.com/sklinkert/go-ddd/internal/domain/entities"
)

func NewDealerResultFromValidatedEntity(dealer *entities.ValidatedDealer) *common.DealerResult {
	return NewDealerResultFromEntity(&dealer.Dealer)
}

func NewDealerResultFromEntity(dealer *entities.Dealer) *common.DealerResult {
	if dealer == nil {
		return nil
	}

	return &common.DealerResult{
		Id:        dealer.Id,
		Name:      dealer.Name,
		CreatedAt: dealer.CreatedAt,
		UpdatedAt: dealer.UpdatedAt,
	}
}
