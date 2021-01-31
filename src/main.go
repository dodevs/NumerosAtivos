package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"sync"
	"time"

	whatsapp "github.com/Rhymen/go-whatsapp"
	"github.com/go-redis/redis/v8"
)

type Block struct {
	Try     func()
	Catch   func(Exception)
	Finally func()
}

type Exception interface{}

func Throw(up Exception) {
	panic(up)
}

func (tcf Block) Do() {
	if tcf.Finally != nil {

		defer tcf.Finally()
	}
	if tcf.Catch != nil {
		defer func() {
			if r := recover(); r != nil {
				tcf.Catch(r)
			}
		}()
	}
	tcf.Try()
}

var countrys = []int{55}
var ddds = []int{27, 28}
var startNumbers = []int{9, 8}

func randomChoice(items []int) interface{} {
	randomIndex := rand.Intn(len(items))
	return items[randomIndex]
}

func generateNumber() string {
	country := randomChoice(countrys)
	ddd := randomChoice(ddds)
	number := fmt.Sprintf("%d%d9%d", country, ddd, randomChoice(startNumbers))
	for i := 1; i <= 7; i++ {
		number += fmt.Sprint(rand.Intn(9))
	}

	number += "@c.us"

	return number
}

func wSession(rdc *redis.Client, c *whatsapp.Conn) {
	rLockAndExecute(
		rdc,
		"whatsapp-session",
		func(r *redis.Client) {
			session := whatsapp.Session{}
			rGet(r, "session", &session)
			if session.ServerToken != "" {
				_, err := c.RestoreWithSession(session)
				errorHandler("restore session", err)
			} else {
				session = wLogin(c)
				rSet(r, "session", session)
			}
		},
	)
}

func verifyNumbers(wcon *whatsapp.Conn, wg *sync.WaitGroup) {
	defer wg.Done()

	if wcon.GetLoggedIn() {

		number := generateNumber()

		for number != "" {

			var exist ExistType
			e, _ := wcon.Exist(number)
			json.Unmarshal([]byte(<-e), &exist)
			if exist.Status == 200 {
				_, err := wcon.SubscribePresence(number)
				errorHandler("get presence", err)
			}
			fmt.Printf("Erro: %v\n", number)

			number = generateNumber()
		}
	}
}

func main() {
	rdc := rConnect()
	//rdc.Del(ctx, "session")
	wcon := wConnect()
	wSession(rdc, wcon)

	csession = cConnect()

	wcon.AddHandler(&waHandler{wcon, uint64(time.Now().Unix())})

	var wg sync.WaitGroup

	for i := 0; i <= 10; i++ {
		wg.Add(1)
		go verifyNumbers(wcon, &wg)
	}

	wg.Wait()
}
