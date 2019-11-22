package utilswebsocket

import (
	"Editor_NaN/EditorApi/models"
)

type kind string
const (
	RUN kind = "run"
	INVITATION = "invitation"
	NEWCLIENT = "newclient"
	ADDTOHUB = "addtohub"
)
func init() {
}


type Message struct {
	sender   uint
	Kind kind `json:"kind"`
	Receivers []uint `json:"rcvs"`
	Users []models.User
	Body interface{} `json:"body"`
}
