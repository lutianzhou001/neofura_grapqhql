package repository

import (
	"context"
	"github.com/lutianzhou001/neo3fura_graphql/graph/model"
	"go.mongodb.org/mongo-driver/bson"
)

func (db *database) Nep11TransfersNotificationsByAddress(address *string) ([]*model.Nep11TransferNotification, error) {
	var results []*model.Nep11TransferNotification
	collection := getCollection(db)
	cursor, err := collection.Find(context.TODO(), bson.M{"$or": []interface{}{
		bson.M{"from": address},
		bson.M{"to": address},
	}})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())
	for cursor.Next(context.TODO()) {
		var v *model.Nep11TransferNotification
		err := cursor.Decode(&v)
		if err != nil {
			return nil, err
		}
		results = append(results, v)
	}
	return results, nil
}

func (db *database) Nep11TransferNotificationByTransactionHash(txid *string) (*model.Nep11TransferNotification, error) {
	var result *model.Nep11TransferNotification
	collection := getCollection(db)
	err := collection.FindOne(context.TODO(), bson.M{"txid": *txid}).Decode(&result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
