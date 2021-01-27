package main

import (
	gocql "github.com/gocql/gocql"
)

func cConnect() *gocql.Session {
	cluster := gocql.NewCluster("127.0.0.1")
	cluster.Keyspace = "numbersdata"
	session, err := cluster.CreateSession()
	errorHandler("gocql create session", err)

	return session
}

func cInsert(session *gocql.Session) {
	// session.Query(
	// 	`INSERT INTO numbers (country, ddd, number, valid, lastView, TODATE(NOW()))`
	// )
}
