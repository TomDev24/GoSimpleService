package main

import (
	"log"
	"github.com/nats-io/stan.go"
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
	sc.Publish("action", []byte("some text"))
}
