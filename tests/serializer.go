package tests

import (
	"github.com/ShuaiXu32167/libsignal-protocol-go/serialize"
)

// newSerializer will return a JSON serializer for testing.
func newSerializer() *serialize.Serializer {
	serializer := serialize.NewJSONSerializer()

	return serializer
}
