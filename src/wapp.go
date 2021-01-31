package main

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	whatsapp "github.com/Rhymen/go-whatsapp"
	"github.com/mitchellh/mapstructure"
	qrcode "github.com/skip2/go-qrcode"
)

type ExistType struct {
	Status int    `json:"status,omitempty`
	Jid    string `json:"jid,omitempty`
}

type Presence struct {
	Deny bool   `json: "deny"`
	ID   string `json: "id"`
	T    int64  `json: "t,omitempty"`
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
	if strings.Contains(data, "Presence") {

		var msg []interface{}
		var presence Presence

		json.Unmarshal([]byte(data), &msg)

		if len(msg) > 1 {
			fmt.Printf("Presence: %v\n", presence)

			mapstructure.Decode(msg[1], &presence)
			cInsert(csession, presence2number(presence))
		}
		fmt.Printf("Erro: %v\n", msg)

	}
}

func wLogin(connection *whatsapp.Conn) whatsapp.Session {
	qrChan := make(chan string)
	go func() {
		// obj := qrcodeTerminal.New()
		// obj.Get(<-qrChan).Print()

		qrcode.WriteFile(<-qrChan, qrcode.Medium, 256, "qr.png")

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
