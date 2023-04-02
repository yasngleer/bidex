package store

import (
	"context"

	"github.com/yasngleer/bidex/types"
)

type UserStore interface {
	Insert(context.Context, *types.User) error
	GetById(context.Context, string) (*types.User, error)
	GetByMail(context.Context, string) (*types.User, error)
}

type SessionStore interface {
	Insert(string, string) error
	Get(string) (string, error)
	Delete(string) error
}

type ItemStore interface {
	Insert(ctx context.Context, items *types.Items) error
	GetById(ctx context.Context, id string) (*types.Items, error)
	GetAll(ctx context.Context) (*[]types.Items, error)
	InsertBid(ctx context.Context, itemid string, bid *types.Bid) error
}
