package utilswebsocket

import (
	// "github.com/astaxie/beego/session"
//	"github.com/gorilla/websocket"
	"fmt"
)

type Store struct {
	name string
	clients  map[uint]*Client
	Stores map[string]*Store
}

func NewStore(name string) *Store {
	s :=Store{
		name: name,
		clients: make(map[uint]*Client),
	}
	return &s
}

func (s *Store) addClient(cl *Client) {
	if oldcl, ok := s.clients[cl.user.Id]; !ok {
		s.clients[cl.user.Id] = cl
		for _, st := range s.Stores {
			if cls := st.getClient(cl.user.Id); cls != nil {
				cl.busy = st
				st.addClient(cl)
			}
		}
	}else {
		oldcl.conn.Close()
		s.clients[cl.user.Id] = cl
	}
}

func (s *Store) addStore(st *Store) {
	if _, ok := s.Stores[st.name]; !ok {
		s.Stores[st.name] = st
	}
}

func (s *Store) removeClient(cl *Client) {
	fmt.Println("remove client", cl.user.Username, "from ", s.name)
	delete(s.clients, cl.user.Id)
}


func (s *Store) removeStore(st *Store) {
	delete(s.Stores, st.name)
}

func (s *Store) getClient(id uint) *Client{
	if oldcl, ok := s.clients[id]; ok {
		return oldcl
	}
	return nil
}

func (s *Store) getStore(name string) *Store{
	if oldcl, ok := s.Stores[name]; ok {
		return oldcl
	}
	return nil
}

