package main

import (
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

func main() {
	// DB initialize
	db := ConnectDB()
	defer db.Disconnect()

	// count
	// count := db.Count()
	// fmt.Println(count)

	start := time.Date(2022, 2, 6, 0, 0, 0, 0, time.UTC)
	end := time.Date(2022, 2, 16, 0, 0, 0, 0, time.UTC)
	step := int64(20)

	var trades []bson.D
	starttime := time.Now()
	db.DistributedQueryByTime(start.UnixMilli(), end.UnixMilli(), step, &trades)
	elapsedtime := time.Since(starttime)

	fmt.Println(len(trades), elapsedtime)
}

func (db *DB)testQuery(date string) {
	// query all
	start := time.Now()

	filter := bson.D{{"td", date}}
	var trades []bson.D
	db.Query(filter, &trades)

	elapsed := time.Since(start)

	fmt.Println(date, len(trades), elapsed)
	// fmt.Println(trades[0])
}
