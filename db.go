package main

import (
        "context"

        "go.mongodb.org/mongo-driver/mongo"
        "go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/bson"
)

const URI = "mongodb://localhost:27017/"

type DB struct {
	client	*mongo.Client
	coll	*mongo.Collection
}

func ConnectDB() DB {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(URI))

        if err != nil {
                panic(err)
        }

	return DB { client, client.Database("trades").Collection("XRP") }
}

func (db *DB) Disconnect() {
	if err := db.client.Disconnect(context.TODO()); err != nil {
		panic(err)
        }
}

func (db *DB) Count() int64 {
	filter := bson.D{}
        count, err := db.coll.CountDocuments(context.TODO(), filter)
        if err != nil {
                panic(err)
        }

	return count
}

func (db *DB) Query(filter, trades interface{}) {
        cursor, err := db.coll.Find(context.TODO(), filter)
        if err != nil {
                panic(err)
        }

        err = cursor.All(context.TODO(), trades)
        if err != nil {
                panic(err)
        }
}
