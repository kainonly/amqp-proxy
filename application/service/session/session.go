package session

import (
	"amqp-proxy/application/service/session/utils"
	"amqp-proxy/application/service/transfer"
	"amqp-proxy/config"
	"errors"
	"github.com/streadway/amqp"
	"go.uber.org/fx"
	"log"
	"time"
)

var (
	QueueNotExists    = errors.New("available queue does not exist")
	QueueIsEmpty      = errors.New("the queue message is empty")
	ReceiptHasExpired = errors.New("the receipt has expired")
	ReceiptIncorrect  = errors.New("the receipt verification is incorrect")
)

type Session struct {
	url             string
	conn            *amqp.Connection
	notifyConnClose chan *amqp.Error
	receipt         *utils.ReceiptMap
	*Dependency
}

type Dependency struct {
	fx.In

	Config   *config.Config
	Transfer *transfer.Transfer
}

type Log struct {
	Queue   string
	Receipt interface{}
	Payload interface{}
	Action  string
}

func New(dep *Dependency) (c *Session, err error) {
	c = new(Session)
	c.Dependency = dep
	c.url = dep.Config.Amqp
	if c.conn, err = amqp.Dial(c.url); err != nil {
		return
	}
	c.notifyConnClose = make(chan *amqp.Error)
	c.conn.NotifyClose(c.notifyConnClose)
	go c.listenConn()
	c.receipt = utils.NewReceiptMap()
	return
}

func (c *Session) listenConn() {
	select {
	case <-c.notifyConnClose:
		log.Println("AMQP connection has been disconnected")
		c.reconnected()
	}
}

func (c *Session) reconnected() {
	var err error
	count := 0
	for {
		time.Sleep(time.Second * 5)
		count++
		log.Println("Trying to reconnect:", count)
		if c.conn, err = amqp.Dial(c.url); err != nil {
			log.Println(err)
			continue
		}
		c.notifyConnClose = make(chan *amqp.Error)
		c.conn.NotifyClose(c.notifyConnClose)
		go c.listenConn()
		log.Println("Attempt to reconnect successfully")
		break
	}
}

func (c *Session) logging(values ...interface{}) error {
	var msg string
	var pipe string
	content := make(map[string]interface{})
	for _, value := range values {
		switch l := value.(type) {
		case PublishOption:
			pipe = c.Config.Transfer.Pipe.Publish
			content = map[string]interface{}{
				"Topic":   l.Exchange,
				"Key":     l.Key,
				"Payload": string(l.Body),
			}
			break
		case Log:
			pipe = c.Config.Transfer.Pipe.Message
			content = map[string]interface{}{
				"Queue":   l.Queue,
				"Receipt": l.Receipt,
				"Payload": l.Payload,
				"Action":  l.Action,
			}
			break
		case error:
			msg = l.Error()
			break
		}
	}
	content["Notice"] = msg
	content["Time"] = time.Now().Unix()
	if pipe != "" {
		return c.Transfer.Push(pipe, content)
	}
	return nil
}
