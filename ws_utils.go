package suprlib

import (
	"golang.org/x/net/websocket"
)

func WebSocketRead(ws *websocket.Conn) (data []byte, err error) {
	var d []byte
	err = websocket.Message.Receive(ws, &d)
	return d, err
}