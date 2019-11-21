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
	kind kind
	sender   *Client
	receivers []*Client
	body interface{}
}
