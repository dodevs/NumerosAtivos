package main

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	whatsapp "github.com/Rhymen/go-whatsapp"
	"github.com/mitchellh/mapstructure"
)

type ExistType struct {
	Status int    `json:"status,omitempty`
	Jid    string `json:"jid,omitempty`
}

type Presence struct {
	Deny bool   `json: "deny"`
	ID   string `json: "id"`
	T    int64  `json: "t"`
	Type string `json: "type"`
}

type waHandler struct {
	wac       *whatsapp.Conn
	startTime uint64
}

func (wh *waHandler) HandleError(err error) {
	errorHandler("wac handldr", err)
}

func (wh *waHandler) HandleJsonMessage(data string) {
	if strings.Contains(data, "Msg") || strings.Contains(data, "Presence") {
		var msg []interface{}
		var presence Presence

		json.Unmarshal([]byte(data), &msg)
		mapstructure.Decode(msg[1], &presence)

		cInsert(csession, presence2number(presence))
	}
}

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
