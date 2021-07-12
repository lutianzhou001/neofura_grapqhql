package repository

import (
	"context"
	"github.com/lutianzhou001/neo3fura_graphql/graph/model"
	"go.mongodb.org/mongo-driver/bson"
)

func (db *database) ScCallsByContractHashAddress(contracthash *string, address *string) ([]*model.ScCall, error) {
	var results []*model.ScCall
	collection := getCollection(db)
	cursor, err := collection.Find(context.TODO(), bson.M{"contractHash": contracthash, "originSender": address})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())
	for cursor.Next(context.TODO()) {
		var v *model.ScCall
		err := cursor.Decode(&v)
		if err != nil {
			return nil, err
		}
		results = append(results, v)
	}
	return results, nil
}
