package p2p

import (
	"fmt"

	"github.com/librity/nc_nomadcoin/utils"
)

type MsgKind int

const (
	MsgLastBlock MsgKind = iota
	MsgAllBlocksReq
	MsgAllBlocksResp
)

type Msg struct {
	Kind    MsgKind
	Payload []byte
}

func (m Msg) String() string {
	kind := m.KindToString()
	str := fmt.Sprintf("{Kind: %s, Payload: []byte{...}}", kind)

	return str
}

func (m *Msg) KindToString() string {
	switch m.Kind {
	case MsgLastBlock:
		return "MsgLastBlock"
	case MsgAllBlocksReq:
		return "MsgAllBlocksReq"
	case MsgAllBlocksResp:
		return "MsgAllBlocksResp"
	default:
		return "UNKOWN"
	}
}

func makeMsgJSON(kind MsgKind, payload interface{}) []byte {
	message := newMsg(kind, payload)
	messageJSON := utils.ToJSON(message)

	return messageJSON
}

func newMsg(kind MsgKind, payload interface{}) *Msg {
	jsonPayload := utils.ToJSON(payload)
	message := &Msg{
		Kind:    kind,
		Payload: jsonPayload,
	}

	return message
}
