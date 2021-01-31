package main

import (
	"os"

	gocql "github.com/gocql/gocql"
)

type Number struct {
	country   int
	ddd       int
	number    string
	valid     bool
	lastView  int64
	updatedAt int64
}

var csession *gocql.Session

func cConnect() *gocql.Session {
	cluster := gocql.NewCluster(os.Getenv("cassandra_host"))
	cluster.Keyspace = "numbersdata"
	session, err := cluster.CreateSession()
	errorHandler("gocql create session", err)

	return session
}

func cInsert(session *gocql.Session, number Number) {
	err := session.Query(
		`INSERT INTO numbers (country, ddd, number, valid, lastView, updatedAt) VALUES (?,?,?,?,?,toTimestamp(NOW()))`,
		number.country, number.ddd, number.number, number.valid, number.lastView).WithContext(ctx).Exec()

	errorHandler("gocql insert", err)
}
