package suprlib

import (
	"golang.org/x/net/websocket"
)

func NewWebsocket(url string, origin string) (*websocket.Conn, error) {
	return websocket.Dial(url, "", origin)
}

func WebsocketRead(conn *websocket.Conn) (data []byte, err error) {
	var d []byte
	err = websocket.Message.Receive(conn, &d)
	return d, err
}

func WebsocketWrite(conn *websocket.Conn, data []byte) error {
	return websocket.Message.Send(conn, data)
}
