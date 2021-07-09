package repository

import (
	"context"
	"github.com/lutianzhou001/neo3fura_graphql/graph/model"
	"go.mongodb.org/mongo-driver/bson"
)

func (db *database) BlockByBlockHash(blockhash *string) (*model.Block, error) {
	var result *model.Block
	collection := getCollection(db)
	err := collection.FindOne(context.TODO(), bson.M{"blockhash": blockhash}).Decode(&result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
