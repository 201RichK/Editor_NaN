package utilswebsocket

import (
	"Editor_NaN/EditorApi/models"
	"github.com/gorilla/websocket"
)

type Client struct {
	user   models.User
	conn *websocket.Conn
	receiver chan *Message
}

func NewClient(c *websocket.Conn, user models.User) *Client {
	client := &Client {
		user: user,
		conn: c,
		receiver: make(chan *Message),
	}
	return client
}