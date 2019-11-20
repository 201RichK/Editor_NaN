package utilswebsocket

import (
	"fmt"
)

func mainHandle(hub *Hub, client *Client) {
	go hub.addClient(client)
	hub.newCompo()
	go 	setInterval(hub.checkConnection, 3)

	comp := hub.Compo["Compo1"]
		for {
			select {
			case m := <- hub.receiver:
				fmt.Println(m)
				switch m.kind {
				case INVITATION:
					for _, each := range m.receivers {
						comp.addClient(client)
						each.receiver <- m
					}
					go eachCompo(comp)
				case NEWCLIENT:
					for _, each := range hub.clients {
						each.receiver <- m
					}
				case ADDTOHUB:
					comp := hub.Compo["Compo1"]					
					comp.addClient(m.sender)
					fmt.Println(comp)
					eachCompo(comp)
				}
			}
		}
}