package utils

import (
	"github.com/streadway/amqp"
	"sync"
)

type ReceiptMap struct {
	sync.RWMutex
	hashMap map[string]*Option
}

type Option struct {
	Queue    string
	Channel  *amqp.Channel
	Delivery *amqp.Delivery
}

func NewReceiptMap() *ReceiptMap {
	c := new(ReceiptMap)
	c.hashMap = make(map[string]*Option)
	return c
}

func (c *ReceiptMap) Put(identity string, receipt *Option) {
	c.Lock()
	c.hashMap[identity] = receipt
	c.Unlock()
}

func (c *ReceiptMap) Empty(identity string) bool {
	return c.hashMap[identity] == nil
}

func (c *ReceiptMap) Get(identity string) *Option {
	c.RLock()
	value := c.hashMap[identity]
	c.RUnlock()
	return value
}

func (c *ReceiptMap) Lists() map[string]*Option {
	return c.hashMap
}

func (c *ReceiptMap) Remove(identity string) {
	delete(c.hashMap, identity)
}
