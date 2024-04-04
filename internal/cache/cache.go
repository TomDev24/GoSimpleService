package cache

import (
	"errors"
	"github.com/TomDev24/GoSimpleService/internal/db"
	"github.com/TomDev24/GoSimpleService/internal/model"
	"log"
)

type Cache struct {
	data map[string]model.Order
	db   *db.Manager
}

func (c *Cache) Init(db *db.Manager) {
	c.data = map[string]model.Order{}
	c.db = db
	c.SetFromDb()
}

func (c *Cache) SetFromDb() error {
	orders, err := c.db.GetAllOrders()
	if err != nil {
		log.Println("Could not fill cache with data from db")
		return err
	}
	for _, o := range orders {
		c.data[o.OrderUid] = o
	}
	return nil
}

func (c *Cache) Save(order model.Order) error {
	_, exist := c.data[order.OrderUid]
	if exist {
		return errors.New("Such order already exist in cache")
	}
	c.data[order.OrderUid] = order
	return nil
}

func (c *Cache) Get(id string) (model.Order, bool) {
	value, ok := c.data[id]
	return value, ok
}
