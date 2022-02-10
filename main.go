package main

import (
	"fmt"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

func main() {
	// DB initialize
	db := ConnectDB()
	defer db.Disconnect()

	filter := bson.D{}
	count, err := db.coll.CountDocuments(context.TODO(), filter)
	if err != nil {
		panic(err)
	}
	fmt.Println(count)

	filter = bson.D{}
	cursor, err := db.coll.Find(context.TODO(), filter)
	if err != nil {
		panic(err)
	}

	start := time.Now()
	var trades []bson.D
	err = cursor.All(context.TODO(), &trades)
	if err != nil {
		panic(err)
	}
	elapsed := time.Since(start)
	fmt.Println(elapsed)

	fmt.Println(trades)
	

}
