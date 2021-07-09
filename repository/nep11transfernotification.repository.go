package repository

import (
	"context"
	"github.com/lutianzhou001/neo3fura_graphql/graph/model"
	"go.mongodb.org/mongo-driver/bson"
)

func (db *database) Nep11TransfersNotificationsByAddress(address *string) ([]*model.Block, error) {
	var result *model.Block
	collection := getCollection(db)
	err := collection.FindOne(context.TODO(), bson.M{"address": *address}).Decode(&result)
	if err != nil {
		return nil, err
	}
	return result, nil
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
