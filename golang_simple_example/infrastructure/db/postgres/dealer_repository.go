package postgres

import (
	"github.com/google/uuid"
	"github.com/sklinkert/go-ddd/internal/domain/entities"
	"github.com/sklinkert/go-ddd/internal/domain/repositories"
	"gorm.io/gorm"
)

type GormDealerRepository struct {
	db *gorm.DB
}

func NewGormDealerRepository(db *gorm.DB) repositories.DealerRepository {
	return &GormDealerRepository{db: db}
}

func (repo *GormDealerRepository) Create(dealer *entities.ValidatedDealer) (*entities.Dealer, error) {
	// Map domain entity to DB model
	dbDealer := toDBDealer(dealer)

	if err := repo.db.Create(dbDealer).Error; err != nil {
		return nil, err
	}

	// Read row from DB to never return different data than persisted
	return repo.FindById(dbDealer.Id)
}

func (repo *GormDealerRepository) FindById(id uuid.UUID) (*entities.Dealer, error) {
	var dbDealer Dealer
	// if err := repo.db.Preload("Seller").First(&dbDealer, id).Error; err != nil {
	// 	return nil, err
	// }

	// Map back to domain entity
	return fromDBDealer(&dbDealer), nil
}

func (repo *GormDealerRepository) FindAll() ([]*entities.Dealer, error) {
	var dbDealers []Dealer

	// if err := repo.db.Preload("").Find(&dbDealers).Error; err != nil {
	// 	return nil, err
	// }

	products := make([]*entities.Dealer, len(dbDealers))
	for i, dbDealer := range dbDealers {
		products[i] = fromDBDealer(&dbDealer)
	}
	return dealers, nil
}

func (repo *GormDealerRepository) Update(dealer *entities.ValidatedDealer) (*entities.Dealer, error) {
	dbDealer := toDBDealer(dealer)
	err := repo.db.Model(&Dealer{}).Where("id = ?", dbDealer.Id).Updates(dbDealer).Error
	if err != nil {
		return nil, err
	}

	// Read row from DB to never return different data than persisted
	return repo.FindById(dbDealer.Id)
}
