package main

import (
	"fmt"
	"time"
	"sync"

	"go.mongodb.org/mongo-driver/bson"
)

func main() {
	// DB initialize
	db := ConnectDB()
	defer db.Disconnect()

	// count
	// count := db.Count()
	// fmt.Println(count)

	var wg sync.WaitGroup

	start := time.Now()

	wg.Add(1)
	go func() {
		db.testQuery("2022-02-07")
		wg.Done()
	}()
	wg.Add(1)
	go func() {
		db.testQuery("2022-02-08")
		wg.Done()
	}()
	wg.Add(1)
	go func() {
		db.testQuery("2022-02-09")
		wg.Done()
	}()
	wg.Add(1)
	go func() {
		db.testQuery("2022-02-10")
		wg.Done()
	}()
	wg.Add(1)
	go func() {
		db.testQuery("2022-02-11")
		wg.Done()
	}()
	wg.Add(1)
	go func() {
		db.testQuery("2022-02-12")
		wg.Done()
	}()

	wg.Wait()

	elapsed := time.Since(start)
	fmt.Println("total:", elapsed)
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
