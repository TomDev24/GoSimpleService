package main

import (
	"os"
	"log"
	"github.com/nats-io/stan.go"
	"time"
	"github.com/brianvoe/gofakeit/v7"
	"encoding/json"
	"github.com/TomDev24/GoSimpleService/internal/model"
	"math/rand/v2"
)

func main() {
	sc, err := stan.Connect("test-cluster", "2", stan.NatsURL("nats://localhost:4222"),
		stan.SetConnectionLostHandler(func(_ stan.Conn, reason error) {
			log.Fatalf("Connection lost, reason: %v", reason)
		}))
	if err != nil{
		log.Fatal(err)
	}
	defer sc.Close()
	for {
		var o model.Order
		if rand.IntN(2) == 1 {
			data, err := os.ReadFile("./msc/model.json")
			if err != nil{
				log.Fatal(err)
			}
			err = json.Unmarshal(data, &o)
		} else {
			err = gofakeit.Struct(&o)
		}
		if err != nil{
			log.Fatal(err)
		}
		b, err := json.Marshal(o)
		if err != nil{
			log.Fatal(err)
		}
		sc.Publish("action", b)
		time.Sleep(3 * time.Second)
	}
}
