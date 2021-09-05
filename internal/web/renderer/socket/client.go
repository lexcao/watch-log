package socket

import (
	"github.com/gorilla/websocket"
	"github.com/lexcao/watch-log/pkg/model"
	log "github.com/sirupsen/logrus"
)

func Sending() chan<- *model.Entry {
	return singleClient.onSending
}

func Receiving() <-chan *model.Entry {
	return singleClient.onReceive
}

var singleClient = &client{}

type client struct {
	connection *websocket.Conn
	onReceive  chan *model.Entry
	onSending  chan *model.Entry
}

func (c *client) ready(conn *websocket.Conn) {
	c.connection = conn
	c.onSending = make(chan *model.Entry)
	c.onReceive = make(chan *model.Entry)
	go c.sending()
	go c.receiving()
}

func (c *client) sending() {
	defer func() {
		c.connection.Close()
	}()

	for {
		select {
		case body, ok := <-c.onSending:
			if !ok {
				c.connection.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			if err := c.connection.WriteJSON(body); err != nil {
				log.Errorf("Sending c error: %v", err)
				return
			}
		}
	}
}

func (c *client) receiving() {
	defer func() {
		c.connection.Close()
		close(c.onSending)
		close(c.onReceive)
	}()

	go func() {
		for body := range c.onReceive {
			log.Info("Receiving", body)
		}
	}()

	for {
		entry := new(model.Entry)
		err := c.connection.ReadJSON(entry)
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Errorf("Receiving c error: %v", err)
			}
			entry.Err = err
		}
		c.onReceive <- entry
	}
}
