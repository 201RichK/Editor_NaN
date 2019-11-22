package utilswebsocket

import (
	"Editor_NaN/EditorApi/models"
	"github.com/gorilla/websocket"
	"encoding/json"
	"fmt"
)

type Client struct {
	user   models.User
	conn *websocket.Conn
	in chan *Message
	out chan *Message
	errc chan error
	busy *Store
}

func NewClient(c *websocket.Conn, user models.User) *Client {
	client := &Client {
		user: user,
		conn: c,
		in: make(chan *Message, 1),
		out: make(chan *Message, 1),
		errc: make(chan error, 1),
	}
	return client
}

func (cl *Client) Handle(s *Store) {
	s.addClient(cl)
	m := Message{
		Kind: "newClient",
	}
	for _, cls := range s.clients {
		m.Users = append(m.Users, cls.user)
	}
	for _, cls := range s.clients {
		cls.out <- &m
	}
	go func() {
		for {
			var m Message
			if err := cl.conn.ReadJSON(&m); err != nil {
				cl.errc <- err
				return
			} else {
				m.sender = cl.user.Id
				cl.in <- &m
			}
		}
	}()



	go func() {

		for m := range cl.out{
			b, err := json.Marshal(*m)
			if err = cl.conn.WriteMessage(1, b); err != nil {
				cl.errc <- err
				return
			}
		}
	}()


	for {
		select {
		case m := <-cl.in:
			
			switch m.Kind {
			case RUN:
				busy := s.getClient(m.sender).busy
				for _, cls := range busy.clients {
					cls.out <- m
				}
			case INVITATION:
				st := NewStore("compo1")
				s.addStore(st)
				sender := s.getClient(m.sender)
				st.addClient(sender)
				sender.busy = st
				for _, id := range m.Receivers {
					if cls, ok := s.clients[id]; ok {
						cls.out <- m
						st.addClient(cls)
						cls.busy = st
					}
				}
				fmt.Println(st)
			}
		case err := <- cl.errc:
			s.removeClient(cl)
			fmt.Println("from errc", err)
			return
		}
	}
}