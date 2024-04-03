package cache

import (
	"github.com/TomDev24/GoSimpleService/internal/model"
)

type Cache struct {
	data  map[string]model.Order
}

func (c *Cache) Init() {
	c.data = map[string]model.Order{}
}

func (c *Cache) Save(order model.Order) {
	c.data[order.OrderUid] = order
}

func (c *Cache) Get(id string) model.Order {
	return c.data[id]
}

