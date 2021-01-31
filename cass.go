package main

import (
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

func cConnect() *gocql.Session {
	cluster := gocql.NewCluster("127.0.0.1")
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
