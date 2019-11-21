package controllers

import (
	"github.com/astaxie/beego"
	"github.com/gorilla/websocket"
	"net/http"
	"Editor_NaN/EditorApi/utils/utilswebsocket"
	"Editor_NaN/EditorApi/models"
	"fmt"
)

var mainHub *utilswebsocket.Hub
func init() {
	mainHub = utilswebsocket.NewHub("mainHUB")
	go utilswebsocket.MainHandle(mainHub)
}

type WebScoketController struct{
	beego.Controller
}

func (this *WebScoketController) Handle() {
	
	ws, err := websocket.Upgrade(this.Ctx.ResponseWriter, this.Ctx.Request, nil, 1024, 1024)
	if _, ok := err.(websocket.HandshakeError); ok {
		http.Error(this.Ctx.ResponseWriter, "Not a websocket handshake", 400)
		return
	} else if err != nil {
		beego.Error("Cannot setup WebSocket connection:", err)
		return
	}
	this.StartSession()
	user := this.GetSession("user").(models.User)
	client := utilswebsocket.NewClient(ws, user)
	fmt.Println("client:", client)
	mainHub.HandleClient(client)
}
