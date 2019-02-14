package protocol

import (
	"github.com/VideoCoin/go-protocol/abis/stream"
)

// Transcoder is a wrapper for the transcoder calls.
type Transcoder struct {
	instance *stream.Stream
	acc      Account
}
