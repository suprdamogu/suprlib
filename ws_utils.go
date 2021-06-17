package suprlib

import (
	"golang.org/x/net/websocket"
)

func NewWebsocket(url string, origin string) (*websocket.Conn, error) {
	return websocket.Dial(url, "", origin)
}

func WebsocketReadBytes(conn *websocket.Conn) (data []byte, err error) {
	var d []byte
	err = websocket.Message.Receive(conn, &d)
	return d, err
}

func WebsocketWriteBytes(conn *websocket.Conn, data []byte) error {
	return websocket.Message.Send(conn, data)
}

func WebsocketReadString(conn *websocket.Conn) (data string, err error) {
	var d string
	err = websocket.Message.Receive(conn, &d)
	return d, err
}

func WebsocketWriteString(conn *websocket.Conn, data string) error {
	return websocket.Message.Send(conn, data)
}
