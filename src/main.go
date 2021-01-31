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

var countrys = []int{55}
var ddds = []int{27, 28}
var startNumbers = []int{9, 8}

type NumberRange struct {
	init int
	end  int
}

var numbersMap = make(map[NumberRange][]NumberRange)

var de27a28 = NumberRange{init: 27, end: 28}
var dddKeys = []NumberRange{de27a28}

func loadNumbers() {
	numbersMap[de27a28] = []NumberRange{{init: 96, end: 99}, {init: 91, end: 94}, {init: 87, end: 88}, {init: 80, end: 83}}
}

func randomChoice(items []NumberRange) NumberRange {
	randomIndex := rand.Intn(len(items))
	return items[randomIndex]
}

func randomRange(start int, end int) int {
	rand.Seed(time.Now().Unix())
	return end + rand.Intn(start-end+1)
}

func gerarNumero() string {
	country := 55

	rangeDDD := randomChoice(dddKeys)
	ddd := randomRange(rangeDDD.init, rangeDDD.end)

	rangeInicial := randomChoice(numbersMap[rangeDDD])
	inicial := randomRange(rangeInicial.init, rangeInicial.end)

	number := fmt.Sprintf("%d%d9%d", country, ddd, inicial)

	for i := 1; i <= 6; i++ {
		number += fmt.Sprint(rand.Intn(9))
	}

	number += "@c.us"

	return number

}

func wSession() {
	rLockAndExecute(
		"whatsapp-session",
		func(r *redis.Client) {
			session := whatsapp.Session{}
			rGet("session", &session)
			if session.ServerToken != "" {
				_, err := wcon.RestoreWithSession(session)
				errorHandler("restore session", err)
			} else {
				session = wLogin()
				rSet("session", session)
			}
		},
	)
}

func verifyNumbers(wcon *whatsapp.Conn, wg *sync.WaitGroup) {
	defer wg.Done()

	if wcon.GetLoggedIn() {

		number := gerarNumero()
		for number != "" {

			var exist ExistType
			e, _ := wcon.Exist(number)
			json.Unmarshal([]byte(<-e), &exist)
			if exist.Status == 200 {
				_, err := wcon.SubscribePresence(number)
				errorHandler("get presence", err)
			} else if exist.Status == 599 {
				break
			}

			number = gerarNumero()
		}
	}
}

func launch() {
	wcon = wConnect()
	wSession()
	wcon.AddHandler(&waHandler{wcon, uint64(time.Now().Unix())})

	var wg sync.WaitGroup

	for i := 0; i <= 0; i++ {
		wg.Add(1)
		go verifyNumbers(wcon, &wg)
	}

	wg.Wait()

	launch()
}

func main() {
	rdc = rConnect()
	//rdc.Del(ctx, "session")
	csession = cConnect()
	launch()
}
