package store

import (
	"context"
	"fmt"

	"github.com/yasngleer/bidex/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoItemStore struct {
	db         *mongo.Database
	collection string
}

func NewMongoItemStore(db *mongo.Database) *MongoItemStore {
	return &MongoItemStore{db: db, collection: "items"}
}

func (s *MongoItemStore) Insert(ctx context.Context, items *types.Items) error {
	if items.Bids == nil {
		items.Bids = &[]types.Bid{}
	}
	objID, _ := primitive.ObjectIDFromHex(items.UserID)
	//make userid objectid instead of string
	bitems, err := bson.Marshal(items)
	if err != nil {
		return err
	}
	bsonitems := bson.M{}
	bson.Unmarshal(bitems, bsonitems)
	bsonitems["user_id"] = objID
	res, err := s.db.Collection(s.collection).InsertOne(ctx, bsonitems)
	if err != nil {
		return err
	}
	items.ID = res.InsertedID.(primitive.ObjectID).Hex()
	return err
}

func (s *MongoItemStore) GetById(ctx context.Context, id string) (*types.Items, error) {
	objID, _ := primitive.ObjectIDFromHex(id)

	qry := bson.A{
		bson.M{
			"$match": bson.M{
				"_id": objID,
			},
		},
		bson.M{
			"$lookup": bson.M{
				"from":         "usersa",
				"localField":   "user_id",
				"foreignField": "_id",
				"as":           "user",
				"pipeline": bson.A{
					bson.M{
						"$project": bson.M{
							"email": 1,
							"id":    1,
						},
					},
				},
			},
		},
		bson.M{
			"$unwind": bson.M{
				"path":                       "$bids",
				"preserveNullAndEmptyArrays": true,
			},
		},
		bson.M{
			"$lookup": bson.M{
				"from":         "users",
				"localField":   "bids.user_id",
				"foreignField": "_id",
				"as":           "bids.user",
				"pipeline": bson.A{
					bson.M{
						"$project": bson.M{
							"email": 1,
							"id":    1,
						},
					},
				},
			},
		},
		bson.M{
			"$group": bson.M{
				"_id": "$_id",
				"name": bson.M{
					"$first": "$name",
				},
				"description": bson.M{
					"$first": "$description",
				},
				"user": bson.M{
					"$first": "$user",
				},
				"image_url": bson.M{
					"$first": "$image_url",
				},
				"bids": bson.M{
					"$push": "$bids",
				},
			},
		},
	}
	cur, err := s.db.Collection(s.collection).Aggregate(ctx, qry)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	item := &types.Items{}
	cur.Next(ctx)
	err = cur.Decode(item)
	if err != nil {
		fmt.Print(err)
		fmt.Print("err")

		return nil, err
	}
	//TODO when there is no bids the query results with empty bid
	bids := *item.Bids
	if bids[0].UserID == "" {
		item.Bids = &[]types.Bid{}
	}
	return item, nil
}

func (s *MongoItemStore) GetAll(ctx context.Context) (*[]types.Items, error) {
	cursor, err := s.db.Collection(s.collection).Find(ctx, bson.D{})
	items := []types.Items{}
	if err != nil {
		defer cursor.Close(ctx)
		return nil, err
	} else {
		for cursor.Next(ctx) {

			item := types.Items{}
			err := cursor.Decode(&item)

			if err != nil {
				return nil, err
			} else {
				items = append(items, item)
			}
		}
	}
	return &items, nil
}

func (s *MongoItemStore) InsertBid(ctx context.Context, itemid string, bid *types.Bid) error {
	objID, _ := primitive.ObjectIDFromHex(itemid)

	objuserID, _ := primitive.ObjectIDFromHex(bid.UserID)

	//make userid objectid instead of string
	bbid, err := bson.Marshal(bid)
	if err != nil {
		return err
	}
	bsonbid := bson.M{}
	bson.Unmarshal(bbid, bsonbid)
	bsonbid["user_id"] = objuserID

	_, err = s.db.Collection(s.collection).UpdateOne(ctx,
		bson.M{"_id": objID},
		bson.M{"$push": bson.M{"bids": bsonbid}},
	)
	if err != nil {
		return err
	}
	return nil

}
