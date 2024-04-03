package stream

import (
	"github.com/nats-io/stan.go"
	"log"
	"encoding/json"
	"github.com/TomDev24/GoSimpleService/internal/model"
	"github.com/TomDev24/GoSimpleService/internal/cache"
	"fmt"
)

type Stream struct {
	scon stan.Conn
	ssub stan.Subscription
	cache *cache.Cache
}

func (s *Stream) handler(msg *stan.Msg) {
	var o model.Order
	err := json.Unmarshal(msg.Data, &o)
	if err != nil{
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", o)
	//d.InsertOrder(o.OrderUid, msg.Data)
	//d.ListAll()
	s.cache.Save(o)
}


func (s *Stream) Init(c *cache.Cache){
	s.cache = c
	sc, err := stan.Connect("test-cluster", "sub", stan.NatsURL("nats://localhost:4222"),
		stan.SetConnectionLostHandler(func(_ stan.Conn, reason error) {
			log.Fatalf("Connection lost, reason: %v", reason)
		}))
	if err != nil{
		log.Fatal(err)
	}
	_, err = sc.Subscribe("action", s.handler)
	if err != nil{
		log.Fatal(err)
	}
	//for {
	//}
}
