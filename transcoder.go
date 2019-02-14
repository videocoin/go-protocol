package protocol

import (
	"github.com/VideoCoin/go-protocol/abis/stream"
)

// TranscoderClient is a wrapper for the transcoder calls.
type TranscoderClient struct {
	instance *stream.Stream
	acc      Account
}
