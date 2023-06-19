package store

import (
	"github.com/yasngleer/bidex/types"
	"gorm.io/gorm"
)

type GormUserStore struct {
	db *gorm.DB
}

func NewGormUserStore(db *gorm.DB) *GormUserStore {
	db.AutoMigrate(types.User{})
	return &GormUserStore{db: db}
}
func (s *GormUserStore) Insert(user *types.User) error {
	err := s.db.Create(user).Error
	return err
}

func (s *GormUserStore) GetById(id int) (*types.User, error) {
	user := &types.User{}
	err := s.db.First(user, types.User{ID: id}).Error
	return user, err
}

func (s *GormUserStore) GetByMail(mail string) (*types.User, error) {
	user := &types.User{}
	err := s.db.First(user, types.User{Email: mail}).Error
	return user, err
}
