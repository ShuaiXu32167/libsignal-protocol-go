package store

import (
	"github.com/ShuaiXu32167/libsignal-protocol-go/groups/state/record"
	"github.com/ShuaiXu32167/libsignal-protocol-go/protocol"
)

type SenderKey interface {
	StoreSenderKey(senderKeyName *protocol.SenderKeyName, keyRecord *record.SenderKey)
	LoadSenderKey(senderKeyName *protocol.SenderKeyName) *record.SenderKey
}
