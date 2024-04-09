package main

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func TestMongo(t *testing.T) {
	t.Run("Test connection to mongo", testMongoConnection)
	t.Run("Test insertion into mongo", testMongoInsertion)
	t.Run("Test reading from mongo", testMongoReading)
}

func testMongoConnection(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017/"))
	assert.NoError(t, err)
	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			t.Fatalf("Error disconnecting from MongoDB: %v", err)
		}
	}()

	err = client.Ping(ctx, nil)
	assert.NoError(t, err)
}

func testMongoInsertion(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017/"))
	assert.NoError(t, err)
	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			t.Fatalf("Error disconnecting from MongoDB: %v", err)
		}
	}()

	collection := client.Database("testdb").Collection("testcollection")

	data := map[string]interface{}{
		"key": "value",
	}

	_, err = collection.InsertOne(ctx, data)
	assert.NoError(t, err)
}

func testMongoReading(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017/"))
	assert.NoError(t, err)
	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			t.Fatalf("Error disconnecting from MongoDB: %v", err)
		}
	}()

	collection := client.Database("testdb").Collection("testcollection")

	cursor, err := collection.Find(ctx, nil)
	if err != nil {
		t.Fatalf("Error finding documents in MongoDB collection: %v", err)
	}
	defer cursor.Close(ctx)

	var results []map[string]interface{}
	for cursor.Next(ctx) {
		var result map[string]interface{}
		err := cursor.Decode(&result)
		assert.NoError(t, err)
		results = append(results, result)
	}

	if len(results) == 0 {
		t.Fatal("No documents found in MongoDB collection")
	}
}