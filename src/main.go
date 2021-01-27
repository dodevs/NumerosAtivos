package main

import (
	"fmt"
	"math/rand"

	"github.com/Rhymen/go-whatsapp"
	"github.com/go-redis/redis/v8"
)

func randomChoice(items []int) interface{} {
	randomIndex := rand.Intn(len(items))
	return items[randomIndex]
}

func generateNumber() string {
	countrys := []int{55}
	ddds := []int{27, 28}
	startNumbers := []int{9, 8}

	number := fmt.Sprintf("%d%d9%d", randomChoice(countrys), randomChoice(ddds), randomChoice(startNumbers))
	for i := 1; i <= 7; i++ {
		number += fmt.Sprint(rand.Intn(9))
	}

	number += "@s.whatsapp.net"

	return number
}

func wConns() *whatsapp.Conn {

}

func wSession(c *whatsapp.Conn) {
	rdc := rConnect()

	rLockAndExecute(
		rdc,
		"whatsapp-session",
		func(r *redis.Client) {
			session := whatsapp.Session{}
			rGet(r, "session", &session)
			if session.ClientId != "" && c.GetConnected() {
				_, err := c.RestoreWithSession(session)
				errorHandler("restore session", err)
			} else {
				session = wLogin(c)
				rSet(r, "session", session)
			}
		},
	)
}

func main() {
	wcon := wConnect()
	wSession(wcon)

	number := generateNumber()
	for number != "" {
		exist, _ := wcon.Exist(number)
		fmt.Printf("%s: %s\n", number, <-exist)
		number = generateNumber()
	}
}
