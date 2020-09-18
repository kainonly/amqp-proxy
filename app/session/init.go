package session

import (
	"amqp-proxy/app/utils"
	"github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
	"time"
)

type Session struct {
	url             string
	conn            *amqp.Connection
	notifyConnClose chan *amqp.Error
	delivery        *utils.SyncDelivery
}

func NewSession(url string) (session *Session, err error) {
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
	session.delivery = utils.NewSyncDelivery()
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
