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

func wConn(rdc *redis.Client) *whatsapp.Conn {
	var wac *whatsapp.Conn
	rLockAndExecute(
		rdc,
		"whatsapp-conn",
		func(r *redis.Client) {
			rGet(r, "conn", &wac)
			if wac == nil {
				wac = wConnect()
				rSet(r, "conn", wac)
			}
		},
	)
	return wac
}

func wSession(rdc *redis.Client, c *whatsapp.Conn) {

	rLockAndExecute(
		rdc,
		"whatsapp-session",
		func(r *redis.Client) {
			session := whatsapp.Session{}
			rGet(r, "session", &session)
			if session.ClientId != "" {
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
	rdc := rConnect()
	//rdc.Del(ctx, "conn")
	wcon := wConn(rdc)
	//wcon := wConnect()
	wSession(rdc, wcon)

	if wcon.GetLoggedIn() {
		number := generateNumber()
		for number != "" {
			exist, _ := wcon.Exist(number)
			fmt.Printf("%s: %s\n", number, <-exist)
			number = generateNumber()
		}
	}

}
