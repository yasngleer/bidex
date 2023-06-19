package store

import (
	"github.com/yasngleer/bidex/types"
)

type UserStore interface {
	Insert(user *types.User) error
	GetById(id int) (*types.User, error)
	GetByMail(mail string) (*types.User, error)
}

type SessionStore interface {
	Insert(string, string) error
	Get(string) (string, error)
	Delete(string) error
}

type ItemStore interface {
	Insert(items *types.Item) error
	GetById(id int) (*types.Item, error)
	GetAll() (*[]types.Item, error)
	InsertBid(bid *types.Bid) error
	GetBidById(id int) (*types.Bid, error)
}
