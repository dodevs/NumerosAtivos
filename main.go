package main

import (
	"fmt"
	"log"
	"time"

	whatsapp "github.com/Rhymen/go-whatsapp"
	gocql "github.com/gocql/gocql"

	goredislib "github.com/go-redis/redis/v8"
	"github.com/go-redsync/redsync/v4"
	"github.com/go-redsync/redsync/v4/redis/goredis/v8"
)

func errorHandler(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func wLogin(connection *whatsapp.Conn) whatsapp.Session {
	qrChan := make(chan string)
	go func() {
		fmt.Printf("qr code: %v\n", <-qrChan)
	}()
	sess, err := connection.Login(qrChan)
	errorHandler(err)

	return sess
}

func wConnect() *whatsapp.Conn {
	wac, err := whatsapp.NewConn(20 * time.Second)
	errorHandler(err)

	return wac
}

func cConnect() *gocql.Session {
	cluster := gocql.NewCluster("127.0.0.1")
	cluster.Keyspace = "numbersdata"
	session, err := cluster.CreateSession()
	errorHandler(err)

	return session
}

func rConnect() *redsync.Redsync {
	/*
		Create a pool with go-redis which is the pool redisync will
		use while communicating with Redis. This can also be any pool that
		implements the 'redis.Pool' interface
	*/
	client := goredislib.NewClient(&goredislib.Options{
		Addr: "localhost:6379",
	})
	pool := goredis.NewPool(client)

	/*
		Create a instance of redisync to be used to obtain a mutual exclusion lock
	*/
	rs := redsync.New(pool)

	/*
		Obtain a new mutex by using the same name for all instances wanting the same lock
	*/
	mutexname := "whatsapp-session"
	mutex := rs.NewMutex(mutexname)

	/*
		Obtain a lock for our given mutex. After this is successful, no one else
		can obtain the same lock (the same mutex name) until we unlock it.
	*/
	err := mutex.Lock()
	errorHandler(err)

	// Do you work that require the lock

	/*
		Release the lock so other processes or threads can obtain a lock
	*/
	_, err = mutex.Unlock()
	errorHandler(err)
}

func main() {
	conn := wConnect()
	_ = wLogin(conn)
}
