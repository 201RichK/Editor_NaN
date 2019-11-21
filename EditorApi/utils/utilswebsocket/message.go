package utilswebsocket

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
	Kind kind `json:"kind"`
	hub *Hub
	sender   *Client
	receivers []*Client
	Body interface{} `json:"body"`
}
