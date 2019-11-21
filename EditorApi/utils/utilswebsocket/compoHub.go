package utilswebsocket

func eachCompo(hub *Hub) {
	for {
		select {
		case m := <- hub.receiver:
			switch m.kind {
			case RUN:
				for _, client := range hub.clients {
					client.receiver <- m
				}
			default:
			}
		}
	}
}