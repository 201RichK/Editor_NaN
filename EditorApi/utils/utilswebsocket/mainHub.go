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
				fmt.Println(m)
				switch m.kind {
				case INVITATION:
					for _, each := range m.receivers {
						each.receiver <- m
					}
				case NEWCLIENT:
					for _, each := range hub.clients {
						each.receiver <- m
					}
				case ADDTOHUB:
				}
			}
		}
}
