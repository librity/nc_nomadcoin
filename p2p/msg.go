package p2p

import (
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
