package session

import (
	"amqp-proxy/app/logging"
	"amqp-proxy/app/types"
	"amqp-proxy/app/utils"
	"github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
	"time"
)

type Session struct {
	url             string
	conn            *amqp.Connection
	notifyConnClose chan *amqp.Error
	receipt         *utils.SyncReceipt
	logging         *logging.Logging
	pipe            *types.PipeOption
}

func NewSession(url string, logging *logging.Logging, pipe *types.PipeOption) (session *Session, err error) {
	session = new(Session)
	session.url = url
	conn, err := amqp.Dial(url)
	if err != nil {
		return
	}
	session.conn = conn
	session.notifyConnClose = make(chan *amqp.Error)
	conn.NotifyClose(session.notifyConnClose)
	go session.listenConn()
	session.receipt = utils.NewSyncReceipt()
	session.logging = logging
	session.pipe = pipe
	return
}

func (c *Session) listenConn() {
	select {
	case <-c.notifyConnClose:
		logrus.Error("AMQP connection has been disconnected")
		c.reconnected()
	}
}

func (c *Session) reconnected() {
	count := 0
	for {
		time.Sleep(time.Second * 5)
		count++
		logrus.Info("Trying to reconnect:", count)
		conn, err := amqp.Dial(c.url)
		if err != nil {
			logrus.Error(err)
			continue
		}
		c.conn = conn
		c.notifyConnClose = make(chan *amqp.Error)
		conn.NotifyClose(c.notifyConnClose)
		go c.listenConn()
		logrus.Info("Attempt to reconnect successfully")
		break
	}
}

func (c *Session) collectFromPublish(option *types.PublishOption, err error) {
	var notice string
	if err != nil {
		notice = err.Error()
	}
	c.logging.Push(c.pipe.Publish, map[string]interface{}{
		"Topic":   option.Exchange,
		"Key":     option.Key,
		"Payload": string(option.Body),
		"Notice":  notice,
		"Time":    time.Now().Unix(),
	})
}

func (c *Session) collectFromAction(queue string, receipt interface{}, payload interface{}, action string, err error) {
	var notice string
	if err != nil {
		notice = err.Error()
	}
	c.logging.Push(c.pipe.Message, map[string]interface{}{
		"Queue":   queue,
		"Receipt": receipt,
		"Payload": payload,
		"Notice":  notice,
		"Action":  action,
		"Time":    time.Now().Unix(),
	})
}
