package main

import (
	"fmt"
	"time"
	"sort"

	"go.mongodb.org/mongo-driver/bson"
)

type Trade struct {
	Tp	float64
	Tv	float64
	Ab	string
	Tms	int64
}

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
	sort.Slice(trades, func(i, j int) bool {
		_i := trades[i].Map()["tms"].(int64)
		_j := trades[j].Map()["tms"].(int64)

		return _i < _j
	})
	elapsedtime := time.Since(starttime)

	fmt.Println(len(trades), elapsedtime)
	fmt.Println(trades[0])
	fmt.Println(trades[len(trades) - 1])
}
