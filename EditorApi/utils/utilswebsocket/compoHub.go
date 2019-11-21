package utilswebsocket

import (
	"fmt"
)

func eachCompo(hub *Hub) {
	for {
		select {
		case m := <- hub.receiver:
			fmt.Println("from compoHUB", m)
			switch m.Kind {
			case RUN:
				fmt.Println("ok")
				for _, client := range hub.clients {
					client.receiver <- m
				}
			default:
			}
		}
	}
}