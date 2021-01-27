package main

import (
	"fmt"
	"time"

	whatsapp "github.com/Rhymen/go-whatsapp"
)

func wLogin(connection *whatsapp.Conn) whatsapp.Session {
	qrChan := make(chan string)
	go func() {
		fmt.Printf("qr code: %v\n", <-qrChan)
	}()
	sess, err := connection.Login(qrChan)
	errorHandler("wapp login", err)

	return sess
}

func wConnect() *whatsapp.Conn {
	wac, err := whatsapp.NewConn(20 * time.Second)
	errorHandler("wapp connect", err)

	return wac
}
