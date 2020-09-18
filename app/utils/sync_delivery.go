package utils

import (
	"github.com/streadway/amqp"
	"sync"
)

type SyncDelivery struct {
	sync.RWMutex
	Map map[string]*amqp.Delivery
}

func NewSyncDelivery() *SyncDelivery {
	c := new(SyncDelivery)
	c.Map = make(map[string]*amqp.Delivery)
	return c
}

func (c *SyncDelivery) Get(identity string) *amqp.Delivery {
	c.RLock()
	value := c.Map[identity]
	c.RUnlock()
	return value
}

func (c *SyncDelivery) Set(identity string, delivery *amqp.Delivery) {
	c.Lock()
	c.Map[identity] = delivery
	c.Unlock()
}
