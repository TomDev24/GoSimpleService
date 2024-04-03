package main

import (
	"github.com/TomDev24/GoSimpleService/internal/db"
	"github.com/TomDev24/GoSimpleService/internal/model"
	"encoding/json"
	"log"
	"github.com/nats-io/stan.go"
)

var d db.Manager;

func handler(msg *stan.Msg) {
	var o model.Order
	err := json.Unmarshal(msg.Data, &o)
	if err != nil{
		log.Fatal(err)
	}
	//fmt.Printf("%+v\n", o)
	d.InsertOrder(o.OrderUid, msg.Data)
	d.ListAll()
}

func main() {
	d.Init()
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
