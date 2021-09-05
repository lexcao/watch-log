package socket

import (
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/lexcao/watch-log/pkg/model"
	log "github.com/sirupsen/logrus"
)

var socket = websocket.Upgrader{
	ReadBufferSize:  4096,
	WriteBufferSize: 4096,
}

type WebSocket struct {
	conn      *websocket.Conn
	OnReceive chan *model.Entry
	ready     chan struct{}
}

func (ws WebSocket) Run(w http.ResponseWriter, r *http.Request) {
	conn, err := socket.Upgrade(w, r, nil)
	if err != nil {
		log.Error(err)
		return
	}

	ws.Ready(conn)
	go ws.Receive()
}

func (ws WebSocket) Send(entry *model.Entry) {
	err := ws.conn.WriteJSON(entry)
	if err != nil {
		log.Warnf("Sending WS message error %s", err)
	}
}

func (ws WebSocket) Receive() {
	entry := new(model.Entry)
	err := ws.conn.ReadJSON(entry)
	if err != nil {
		log.Warnf("Receiving WS message error %s", err)
	}
	log.Infof("Receiving WS message: %v", entry)
	ws.OnReceive <- entry
}

func (ws WebSocket) IsReady() <-chan struct{} {
	return ws.ready
}

func (ws WebSocket) Ready(conn *websocket.Conn) {
	ws.conn = conn
	ws.ready <- struct{}{}
}
