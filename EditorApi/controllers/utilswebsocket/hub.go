package utilswebsocket

import (
	"github.com/astaxie/beego/session"
//	"github.com/gorilla/websocket"
)

type Hub struct {
	name string
	clients  map[uint]*Client
	receiver chan *Message
	Session  session.Store
	Compo map[string]*Hub
}


// Creer un nouveau Hub
func newHub(name string) *Hub {
	return &Hub {
		name: name,
		clients: make(map[uint]*Client),
		Compo: make(map[string]*Hub),
		receiver: make(chan *Message),
	}
}


func (hub *Hub) newCompo() {
	hub.Compo["Compo1"] = newHub("Compo1")
}



// Ajouter un client au Hub specifique
func (hub *Hub) addClient(client *Client) {
	if client != nil {
		if _, ok := hub.clients[client.user.Id]; !ok {
			hub.clients[client.user.Id] = client
		} else {
			hub.deleteClient(client.user.Id)
			hub.clients[client.user.Id] = client
		}
	}
	go hub.handleEachClient()
}



// Obtenir un Client par son id
func (hub *Hub) getClient(id uint) *Client {
	if client, ok := hub.clients[id]; ok {
		return client
	}
	return nil
}



// Supprimer et deconnecter un client
func (hub *Hub) deleteClient(id uint) {
	if client, ok := hub.clients[id]; ok {
		delete(hub.clients, id)
		client.conn.Close()
	}
}


// Supprimer un client de la liste du hub specifique
func (hub *Hub) removeClient(id uint) {
	delete(hub.clients, id)
}



// verifier si les clients sont toujours connectés
func (hub *Hub) checkConnection() {
	for _, client := range hub.clients {
		if err := client.conn.ReadJSON(nil); err != nil {
			hub.deleteClient(client.user.Id)
		}
	}
}




func (hub * Hub) handleEachClient() {
	for _, client := range hub.clients {
		go func () {
				for {
					m := <-client.receiver
					client.conn.WriteJSON(m.body)
				}

		}()
		for {
			m := new(Message)
			m.kind = ADDTOHUB
			client.conn.ReadJSON(&m.body)
			m.sender = client

			hub.receiver <- m
		}
	}
}



func run(handle func()) {
	handle()
}
