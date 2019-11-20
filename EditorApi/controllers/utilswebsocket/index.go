package utilswebsocket

import (
	"Editor_NaN/EditorApi/models"
	"github.com/astaxie/beego/session"
	"github.com/gorilla/websocket"
	"time"
)


var mainHub *Hub
func init() {
	mainHub = newHub("mainHUB")
}

func Handle(c *websocket.Conn, ss session.Store) {
	mainHandle(mainHub, NewClient(c, ss.Get("user").(models.User)))
}


func setInterval(someFunc func(), milliseconds int) {
	interval := time.Duration(milliseconds) * time.Second
	ticker := time.NewTicker(interval)

	for {
		select {
		case <-ticker.C:
			go someFunc()
		}
	}
}