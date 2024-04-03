package main

import (
	"fmt"
	"log"
	"github.com/nats-io/stan.go"
)

func handler(msg *stan.Msg) {
	fmt.Println("msg:", msg)
}

func main() {
	//sc, err := stan.Connect("test-cluster", "1", )
	sc, err := stan.Connect("test-cluster", "1", stan.NatsURL("nats://localhost:4222"), 
		stan.SetConnectionLostHandler(func(_ stan.Conn, reason error) {
			log.Fatalf("Connection lost, reason: %v", reason)
		}))
	if err != nil{
		log.Fatal(err)
	}
	_, err = sc.Subscribe("action", handler)
	if err != nil{
		log.Fatal(err)
	}
	for {
	}
}
