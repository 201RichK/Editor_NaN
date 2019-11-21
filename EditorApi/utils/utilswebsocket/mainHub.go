package utilswebsocket

import (
	"fmt"
)





// Traitement de tous les clients dans le Hub principal
func MainHandle(hub *Hub) {
	go setInterval(hub.checkConnection, 3)
		for {
			select {
			case m := <- hub.receiver: 
				fmt.Println("from mainHUB", m)
				switch m.Kind {
				case INVITATION:
					for _, each := range m.receivers {
						each.receiver <- m
					}
					
				case NEWCLIENT:
					for _, each := range hub.clients {
						each.receiver <- m
					}
				case ADDTOHUB:
					compoHUB := hub.Compo["Compo1"]
					hub.removeClient(m.sender.user.Id)
					compoHUB.HandleClient(m.sender)
					fmt.Println(compoHUB)
				case RUN:
					for _, client := range m.hub.clients {
						if client != m.sender {
							client.receiver <- m
							fmt.Println("each", client)
						}
					}
				}
			}
		}
}
