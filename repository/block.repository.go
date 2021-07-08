package repository

import (
	"context"
	"fmt"
	"github.com/lutianzhou001/neo3fura_graphql/graph/model"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type BlockRepository interface {
	FindAll() []*model.Block
	FindLastBlock() *model.Block
}

type database struct {
	client *mongo.Client
}

func NewBlockRepository() BlockRepository {
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
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	return &database{
		client: userClient,
	}
}

func (db *database) FindLastBlock() *model.Block {
	var result *model.Block
	collection := db.client.Database("neo").Collection("Block")
	err := collection.FindOne(context.TODO(),bson.M{}).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}
	return result
}

func (db *database) FindAll() []*model.Block {
	collection := db.client.Database("neo").Collection("Block")
	cursor, err := collection.Find(context.TODO(), bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(context.TODO())
	var results []*model.Block
	for cursor.Next(context.TODO()) {
		var v *model.Block
		err := cursor.Decode(&v)
		if err != nil {
			log.Fatal(err)
		}
		results = append(results, v)
	}
	return results
}
