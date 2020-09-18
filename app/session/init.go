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
	channel         *utils.SyncChannel
	channelDone     *utils.SyncChannelDone
	channelReady    *utils.SyncChannelReady
	notifyChanClose *utils.SyncNotifyChanClose
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
	session.channel = utils.NewSyncChannel()
	session.channelDone = utils.NewSyncChannelDone()
	session.channelReady = utils.NewSyncChannelReady()
	session.notifyChanClose = utils.NewSyncNotifyChanClose()
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

func (c *Session) SetChannel(ID string) (err error) {
	var channel *amqp.Channel
	channel, err = c.conn.Channel()
	if err != nil {
		return
	}
	c.channel.Set(ID, channel)
	c.channelDone.Set(ID, make(chan int))
	notifyChanClose := make(chan *amqp.Error)
	channel.NotifyClose(notifyChanClose)
	c.notifyChanClose.Set(ID, notifyChanClose)
	go c.listenChannel(ID)
	return
}

func (c *Session) listenChannel(ID string) {
	select {
	case <-c.notifyChanClose.Get(ID):
		logrus.Error("Channel connection is disconnected:", ID)
		if c.channelReady.Get(ID) {
			c.refreshChannel(ID)
		} else {
			break
		}
	case <-c.channelDone.Get(ID):
		break
	}
}

func (c *Session) refreshChannel(ID string) {
	for {
		err := c.SetChannel(ID)
		if err != nil {
			continue
		}
		logrus.Info("Channel refresh successfully")
		break
	}
}
