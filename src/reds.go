package main

import (
	"bytes"
	"context"
	"encoding/gob"
	"fmt"

	"github.com/go-redis/redis/v8"
	goredislib "github.com/go-redis/redis/v8"
	"github.com/go-redsync/redsync/v4"
	"github.com/go-redsync/redsync/v4/redis/goredis/v8"
)

type call func(rdc *goredislib.Client)

var ctx = context.Background()
var rdc *redis.Client

func rConnect() *goredislib.Client {
	client := goredislib.NewClient(&goredislib.Options{
		Addr:     fmt.Sprintf("%s:6379", "127.0.0.1"),
		Password: "fh7gdGDds34",
	})

	return client
}

func rLockAndExecute(mtxname string, fn call) {
	pool := goredis.NewPool(rdc)
	rs := redsync.New(pool)

	mutex := rs.NewMutex(mtxname)
	err := mutex.Lock()
	errorHandler("mutex lock", err)

	fn(rdc)

	_, err = mutex.Unlock()
	errorHandler("mutex unlok", err)
}

func rSet(key string, value interface{}) {
	encoded := new(bytes.Buffer)
	err := gob.NewEncoder(encoded).Encode(value)
	errorHandler("gob encode", err)

	_, err = rdc.Do(ctx, "SET", key, encoded.Bytes()).Result()
	errorHandler("redis hset", err)
}

func rGet(key string, dest interface{}) {
	res, err := rdc.Do(ctx, "GET", key).Result()
	errorHandler("redis get", err)

	if res != nil {
		encoded := []byte(res.(string))
		encodedReceivedIOReader := bytes.NewBuffer(encoded)
		err = gob.NewDecoder(encodedReceivedIOReader).Decode(dest)
		errorHandler("redis hget", err)
	}

}
