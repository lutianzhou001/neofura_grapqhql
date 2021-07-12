package repository

import (
	"context"
	"github.com/lutianzhou001/neo3fura_graphql/graph/model"
	"go.mongodb.org/mongo-driver/bson"
)

func (db *database) ScVoteCallByTransactionHash(transactionhash *string) ([]*model.ScVoteCall, error) {
	var result []*model.ScVoteCall
	collection := getCollection(db)
	err := collection.FindOne(context.TODO(), bson.M{"txid": transactionhash}).Decode(&result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (db *database) ScVoteCallsByVoterAddress(voteraddress *string) ([]*model.ScVoteCall, error) {
	var results []*model.ScVoteCall
	collection := getCollection(db)
	cursor, err := collection.Find(context.TODO(), bson.M{"voter": voteraddress})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())
	for cursor.Next(context.TODO()) {
		var v *model.ScVoteCall
		err := cursor.Decode(&v)
		if err != nil {
			return nil, err
		}
		results = append(results, v)
	}
	return results, nil
}

func (db *database) ScVoteCallsByCandidateAddress(candidateaddress *string) ([]*model.ScVoteCall, error) {
	var results []*model.ScVoteCall
	collection := getCollection(db)
	cursor, err := collection.Find(context.TODO(), bson.M{"candidate": candidateaddress})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())
	for cursor.Next(context.TODO()) {
		var v *model.ScVoteCall
		err := cursor.Decode(&v)
		if err != nil {
			return nil, err
		}
		results = append(results, v)
	}
	return results, nil
}
