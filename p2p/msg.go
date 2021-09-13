package p2p

type MsgKind int

const (
	LastBlock MsgKind = iota
	MsgAllBlocksRequest
	MsgAllBlocksResponse
)

type Msg struct {
	Kind    MsgKind
	Payload []byte
}

func newMsg() {

}
