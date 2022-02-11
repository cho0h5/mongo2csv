package main

import (
        "context"
	"sync"

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

func (db *DB) QueryByTime(start, end int64, trades interface{}) {
	filter := bson.M{"tms": bson.M{"$gte": start, "$lt": end}}
	db.Query(filter, trades)
}

func (db *DB) DistributedQueryByTime(start, end int64, step int64, trades *[]bson.D) {
	gap := (end - start) / step

	var wg sync.WaitGroup
	wg.Add(int(step))

	result := make(chan []bson.D)

	for i := int64(0); i < step; i++ {
		_start := start + i * gap
		_end := _start + gap
		go func(i, start, end int64) {
			defer wg.Done()

			var trades []bson.D

			db.QueryByTime(start, end, &trades)

			result <- trades
		}(i, _start, _end)
	}

	var wg_reducer sync.WaitGroup
	wg_reducer.Add(1)

	go func() {
		defer wg_reducer.Done()
		for trade := range result {
			for _, value := range trade {
				*trades = append(*trades, value)
			}
		}
	}()

	wg.Wait()
	close(result)

	wg_reducer.Wait()
}
