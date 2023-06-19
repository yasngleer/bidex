package store

import (
	"github.com/yasngleer/bidex/types"
	"gorm.io/gorm"
)

type GormItemStore struct {
	db *gorm.DB
}

func NewGormItemStore(db *gorm.DB) *GormItemStore {
	db.AutoMigrate(types.Item{})
	db.AutoMigrate(types.Bid{})

	return &GormItemStore{db: db}
}

func (s *GormItemStore) Insert(items *types.Item) error {
	err := s.db.Create(items).Error
	return err
}

func (s *GormItemStore) GetById(id int) (*types.Item, error) {
	item := types.Item{}
	err := s.db.Where(&types.Item{ID: id}).
		Preload("User").
		Preload("Bids").
		Preload("Bids.User").
		Find(&item).Error
	return &item, err
}

func (s *GormItemStore) GetAll() (*[]types.Item, error) {
	todos := []types.Item{}
	err := s.db.Where(types.Item{}).Preload("User").Preload("Bids").Find(&todos).Error
	return &todos, err
}

func (s *GormItemStore) InsertBid(bid *types.Bid) error {
	return s.db.Create(bid).Error
}
func (s *GormItemStore) GetBidById(id int) (*types.Bid, error) {
	bid := types.Bid{}
	err := s.db.Where(types.Bid{ID: id}).Preload("User").Find(&bid).Error
	return &bid, err
}
