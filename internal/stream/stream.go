package stream

import (
	"encoding/json"
	"github.com/TomDev24/GoSimpleService/internal/cache"
	"github.com/TomDev24/GoSimpleService/internal/db"
	"github.com/TomDev24/GoSimpleService/internal/model"
	"github.com/go-playground/validator/v10"
	"github.com/nats-io/stan.go"
	"log"
)

type Stream struct {
	scon     stan.Conn
	ssub     stan.Subscription
	cache    *cache.Cache
	db       *db.Manager
	validate *validator.Validate
}

func (s *Stream) Init(c *cache.Cache, db *db.Manager) error {
	sc, err := stan.Connect("test-cluster", "sub", stan.NatsURL("nats://localhost:4222"),
		stan.SetConnectionLostHandler(func(_ stan.Conn, reason error) {
			log.Fatalf("Connection lost, reason: %v", reason)
		}))
	if err != nil {
		return err
	}
	sub, err := sc.Subscribe("NewOrder", s.handler)
	if err != nil {
		return err
	}
	s.cache, s.db = c, db
	s.scon, s.ssub = sc, sub
	s.validate = validator.New()
	return nil
}

func (s *Stream) Close() {
	s.ssub.Unsubscribe()
	s.scon.Close()
}

func (s *Stream) handler(msg *stan.Msg) {
	var o model.Order

	log.Println("Incoming data from channel:", msg)
	err := json.Unmarshal(msg.Data, &o)
	if err != nil {
		log.Println(err)
		return
	}
	err = s.validate.Struct(o)
	if err != nil {
		log.Println("Skiping order: validation error ", err)
		return
	}
	_, exist := s.cache.Get(o.OrderUid)
	if exist {
		log.Println("Skiping order: such order already exist")
		return
	}
	err = s.db.InsertOrder(o.OrderUid, msg.Data)
	if err != nil {
		log.Println("Could not insert into db", err)
		return
	}
	err = s.cache.Save(o)
	if err != nil {
		log.Println(err)
	}
}
