package protocol

import (
	"github.com/VideoCoin/go-protocol/abis/stream"
)

// VerifierClient is a wrapper for the transcoder calls.
type VerifierClient struct {
	instance *stream.Stream
	caller   Caller
}
