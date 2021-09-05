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
	port int
}

func NewWebRenderer(port int) component.Renderer {
	renderer := &WebRenderer{
		port: port,
	}

	go renderer.initServer()

	log.Infof("Watch Log start to serve logs to web, listening on http://localhost:%d", port)

	return renderer
}

func (r *WebRenderer) Render(entry *model.Entry) {
	if len(entry.PipelinedObject) == 0 {
		return
	}
	select {
	case socket.Sending() <- entry:
	case <-time.After(1 * time.Second):
		log.Info("WebSocket is not ready, pending for connection")
		r.Render(entry)
	}
}

func (r *WebRenderer) initServer() {
	http.Handle("/", http.FileServer(http.Dir("./templates")))
	http.HandleFunc("/ws", socket.WebSocket())
	err := http.ListenAndServe(fmt.Sprintf("localhost:%d", r.port), nil)
	if err != nil {
		log.Fatal(err)
	}
}
