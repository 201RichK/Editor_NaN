package utilswebsocket

import (
	"github.com/astaxie/beego/session"
//	"github.com/gorilla/websocket"
"fmt"
)

type Hub struct {
	name string
	clients  map[uint]*Client
	receiver chan *Message
	Session  session.Store
	Compo map[string]*Hub
}


// Creer un nouveau Hub
func NewHub(name string) *Hub {
	return &Hub {
		name: name,
		clients: make(map[uint]*Client),
		Compo: make(map[string]*Hub),
		receiver: make(chan *Message),
	}
}


func (hub *Hub) newCompo() {
	hub.Compo["Compo1"] = NewHub("Compo1")
}



// Ajouter un client au Hub specifique
func (hub *Hub) addClient(client *Client) {
		if _, ok := hub.clients[client.user.Id]; !ok {
			hub.clients[client.user.Id] = client
		}else {
			hub.deleteClient(client.user.Id)
			hub.clients[client.user.Id] = client
		}
		fmt.Println(hub)
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
		fmt.Println("connction close")
	}
}


// Supprimer un client de la liste du hub specifique
func (hub *Hub) removeClient(id uint) {
	delete(hub.clients, id)
}



// verifier si les clients sont toujours connectés
func (hub *Hub) checkConnection() {
	for _, client := range hub.clients {
		if msg, _, err := client.conn.ReadMessage(); err != nil {
			if msg != -1 {
				hub.deleteClient(client.user.Id)
			}
		}
	}
}




func (hub * Hub) listensClient(client *Client) {
		go func () {
				for {
					m := <-client.receiver
					client.conn.WriteJSON(m.body)
				}

		}()
		//for {
		//	m := new(Message)
		// 	m.kind = ADDTOHUB
		// 	err := client.conn.ReadJSON(&m.body)
		// 	if err != nil {
		// 		panic(err)
		// 	} else {
		// 		m.sender = client
		// 		hub.receiver <- m
		// 	}
		// }
}






func (hub *Hub) HandleClient(client *Client) {
	fmt.Println("my hub", hub)

	hub.addClient(client)
	go hub.listensClient(client)
}
