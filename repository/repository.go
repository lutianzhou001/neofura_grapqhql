package repository

import (
	"context"
	"fmt"
	"github.com/lutianzhou001/neo3fura_graphql/graph/model"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type BlockRepository interface {
	BlockByBlockHash(blockhash *string) (*model.Block, error)
}

type Nep11TransferNotificationRepository interface {
	Nep11TransfersNotificationsByAddress(address *string) ([]*model.Nep11TransferNotification, error)
	Nep11TransferNotificationByTransactionHash(txid *string) (*model.Nep11TransferNotification, error)
}

type ScCallRepository interface {
	ScCallsByContractHashAddress(contracthash *string, address *string) ([]*model.ScCall, error)
}

type ScVoteCallRepository interface {
	ScVoteCallByTransactionHash(transactionhash *string) ([]*model.ScVoteCall, error)
	ScVoteCallsByVoterAddress(voteraddress *string) ([]*model.ScVoteCall, error)
	ScVoteCallsByCandidateAddress(candidateaddress *string) ([]*model.ScVoteCall, error)
}

type database struct {
	client *mongo.Client
}

func getCollection(db *database) *mongo.Collection {
	return db.client.Database("neo").Collection("Block")
}

func getConnection() (uc *mongo.Client, err error) {
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	// mongodb+srv://USERNAME:PASSWORD@HOST:PORT
	// MONGODB := os.Getenv("MONGODB")
	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb://neo:pIqGZYY@40.76.139.118:27018/neo")
	clientOptions = clientOptions.SetMaxPoolSize(50)
	// Connect to MongoDB
	userClient, err := mongo.Connect(ctx, clientOptions)

	if err != nil {
		log.Fatal(err)
	}
	// Check the connection
	err = userClient.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}
	fmt.Println("Connected to MongoDB!")
	return userClient, nil
}

func NewBlockRepository() BlockRepository {
	uc, err := getConnection()
	if err != nil {
		log.Fatal(err)
	}
	return &database{
		client: uc,
	}
}

func NewNep11TransferNotificationRepository() Nep11TransferNotificationRepository {
	uc, err := getConnection()
	if err != nil {
		log.Fatal(err)
	}
	return &database{
		client: uc,
	}
}

func NewScVoteCallRepository() ScVoteCallRepository {
	uc, err := getConnection()
	if err != nil {
		log.Fatal(err)
	}
	return &database{
		client: uc,
	}
}

func NewScCallRepository() ScCallRepository {
	uc, err := getConnection()
	if err != nil {
		log.Fatal(err)
	}
	return &database{
		client: uc,
	}
}