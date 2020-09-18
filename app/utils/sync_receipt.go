package utils

import (
	"amqp-proxy/app/types"
	"sync"
)

type SyncReceipt struct {
	sync.RWMutex
	Map map[string]*types.ReceiptOption
}

func NewSyncReceipt() *SyncReceipt {
	c := new(SyncReceipt)
	c.Map = make(map[string]*types.ReceiptOption)
	return c
}

func (c *SyncReceipt) Get(identity string) *types.ReceiptOption {
	c.RLock()
	value := c.Map[identity]
	c.RUnlock()
	return value
}

func (c *SyncReceipt) Set(identity string, receipt *types.ReceiptOption) {
	c.Lock()
	c.Map[identity] = receipt
	c.Unlock()
}

func (c *SyncReceipt) Delete(identity string) {
	delete(c.Map, identity)
}
