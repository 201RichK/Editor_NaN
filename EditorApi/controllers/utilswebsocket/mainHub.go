package utilswebsocket

import (
	"fmt"
)

func mainHandle(hub *Hub, client *Client) {
	go hub.addClient(client)
	hub.newCompo()
	comp := hub.Compo["Compo1"]
	go func () {
		for {
			select {
			case m := <- hub.receiver:
				switch m.kind {
				case INVITATION:
					for _, each := range m.receivers {
						comp.addClient(client)
						each.receiver <- m
					}
					fmt.Println(comp)
					go eachCompo(comp)
				case NEWCLIENT:
					for _, each := range hub.clients {
						each.receiver <- m
					}
				case ADDTOHUB:
					comp := hub.Compo["Compo1"]					
					comp.addClient(m.sender)
					eachCompo(comp)
					fmt.Println(comp)

				}
			}
		}
	}()
	setInterval(hub.checkConnection, 3)
}