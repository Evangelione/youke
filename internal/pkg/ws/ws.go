package ws

import (
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

var Socket websocket.Upgrader

func init() {
	Socket = websocket.Upgrader{
		HandshakeTimeout: 5 * time.Second,
		ReadBufferSize:   1024,
		WriteBufferSize:  1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
}
