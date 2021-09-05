package renderer

import (
	"fmt"
	"net/http"
	"time"

	"github.com/lexcao/watch-log/internal/web/renderer/socket"
	"github.com/lexcao/watch-log/pkg/component"
	"github.com/lexcao/watch-log/pkg/model"
	log "github.com/sirupsen/logrus"
)

type WebRenderer struct {
	ws   *socket.WebSocket
	port int
}

func NewWebRenderer(port int) component.Renderer {
	renderer := &WebRenderer{
		ws:   &socket.WebSocket{},
		port: port,
	}

	go renderer.initServer()

	log.Infof("Watch Log start to serve logs to web, listening on http://localhost:%d", port)

	return renderer
}

func (r *WebRenderer) Render(entry *model.Entry) {
	select {
	case <-r.ws.IsReady():
		r.ws.Send(entry)
	case <-time.After(1 * time.Second):
		log.Info("WebSocket is not ready, pending for connection")
		r.Render(entry)
	}
}

func (r *WebRenderer) initServer() {
	port := r.port
	client := r.ws
	http.Handle("/", http.FileServer(http.Dir("./templates")))
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		client.Run(w, r)
	})
	err := http.ListenAndServe(fmt.Sprintf("localhost:%d", port), nil)
	if err != nil {
		log.Fatal(err)
	}
}
