package store

import (
	"context"

	"github.com/yasngleer/bidex/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoUserStore struct {
	db     *mongo.Database
	column string
}

func NewMongoUserStore(db *mongo.Database) *MongoUserStore {
	return &MongoUserStore{
		db:     db,
		column: "users",
	}
}

func (s *MongoUserStore) Insert(ctx context.Context, u *types.User) error {
	res, err := s.db.Collection(s.column).InsertOne(ctx, u)
	if err != nil {
		return err
	}
	u.ID = res.InsertedID.(primitive.ObjectID).Hex()
	return err
}

func (s *MongoUserStore) GetById(ctx context.Context, id string) (*types.User, error) {
	var (
		objID, _ = primitive.ObjectIDFromHex(id)
		res      = s.db.Collection(s.column).FindOne(ctx, bson.M{"_id": objID})
		u        = &types.User{}
		err      = res.Decode(u)
	)
	return u, err
}
func (s *MongoUserStore) GetByMail(ctx context.Context, mail string) (*types.User, error) {
	var (
		res = s.db.Collection(s.column).FindOne(ctx, bson.M{"email": mail})
		u   = &types.User{}
		err = res.Decode(u)
	)
	return u, err
}
