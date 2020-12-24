package ZZWebsocket

import (
	"golang.org/x/net/websocket"
)

func WSRead(ws *websocket.Conn) (data []byte, err error) {
	var d []byte
	err = websocket.Message.Receive(ws, &d)
	return d, err
}
