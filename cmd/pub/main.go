package main

import (
	"encoding/json"
	"github.com/TomDev24/GoSimpleService/internal/model"
	"github.com/brianvoe/gofakeit/v7"
	"github.com/nats-io/stan.go"
	"log"
	"math/rand/v2"
	"os"
	"time"
)

type Bar struct {
	Name   string
	Number int
	Float  float32
}

func main() {
	var o model.Order
	var bar Bar
	var bytes []byte

	sc, err := stan.Connect("test-cluster", "pub", stan.NatsURL("nats://localhost:4222"),
		stan.SetConnectionLostHandler(func(_ stan.Conn, reason error) {
			log.Fatalf("Connection lost, reason: %v", reason)
		}))
	if err != nil {
		log.Fatal(err)
	}
	defer sc.Close()

	for {
		switch dice := rand.IntN(3); dice {
		case 0:
			// send correct model
			bytes, err = os.ReadFile("./msc/model.json")
			if err != nil {
				log.Fatal(err)
			}
		case 1:
			// send correct model, with fake data
			err = gofakeit.Struct(&o)
			if err != nil {
				log.Fatal(err)
			}
			bytes, err = json.Marshal(o)
		case 2:
			// send incorrect model
			err = gofakeit.Struct(&bar)
			if err != nil {
				log.Fatal(err)
			}
			bytes, err = json.Marshal(&bar)
		}
		if err != nil {
			log.Fatal(err)
		}
		sc.Publish("NewOrder", bytes)
		time.Sleep(3 * time.Second)
	}
}
