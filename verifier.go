package protocol

import (
	"github.com/VideoCoin/go-protocol/abis/stream"
)

// Verifier is a wrapper for the transcoder calls.
type Verifier struct {
	instance *stream.Stream
	acc      Account
}
