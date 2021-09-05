package socket

import (
	"net/http"
	"strings"

	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
)

var socket = websocket.Upgrader{
	ReadBufferSize:  4096,
	WriteBufferSize: 4096,
	CheckOrigin: func(r *http.Request) bool {
		return strings.Contains(r.Host, "localhost")
	},
}

func WebSocket() http.HandlerFunc {
	singleClient = &client{}

	return func(w http.ResponseWriter, r *http.Request) {
		conn, err := socket.Upgrade(w, r, nil)
		if err != nil {
			log.Error(err)
			return
		}

		singleClient.ready(conn)
	}
}
