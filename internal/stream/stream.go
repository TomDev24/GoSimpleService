package stream

import (
	"encoding/json"
	"log"
	"github.com/TomDev24/GoSimpleService/internal/cache"
	"github.com/TomDev24/GoSimpleService/internal/db"
	"github.com/TomDev24/GoSimpleService/internal/model"
	"github.com/nats-io/stan.go"
)

type Stream struct {
	scon stan.Conn
	ssub stan.Subscription
	cache *cache.Cache
	db		*db.Manager
}

func (s *Stream) Init(c *cache.Cache, db *db.Manager) error {
	sc, err := stan.Connect("test-cluster", "sub", stan.NatsURL("nats://localhost:4222"),
		stan.SetConnectionLostHandler(func(_ stan.Conn, reason error) {
			log.Fatalf("Connection lost, reason: %v", reason)
		}))
	if err != nil{
		return err
	}
	sub, err := sc.Subscribe("NewOrder", s.handler)
	if err != nil{
		return err
	}
	s.cache, s.db = c, db
	s.scon, s.ssub = sc, sub
	return nil
}

func (s *Stream) Close() {
	s.ssub.Unsubscribe()
	s.scon.Close()
}

func (s *Stream) handler(msg *stan.Msg) {
	var o model.Order

	err := json.Unmarshal(msg.Data, &o)
	if err != nil{
		log.Println(err)
	}
	err = s.db.InsertOrder(o.OrderUid, msg.Data)
	if err != nil{
		log.Println(err)
		return
	}
	err = s.cache.Save(o)
	if err != nil{
		log.Println(err)
	}
}
