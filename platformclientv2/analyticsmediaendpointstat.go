package platformclientv2
import (
	"encoding/json"
)

// Analyticsmediaendpointstat
type Analyticsmediaendpointstat struct { 
	// Codecs - The MIME types of the audio encodings used by the audio streams belonging to this endpoint
	Codecs *[]string `json:"codecs,omitempty"`


	// MinMos - The lowest estimated average MOS among all the audio streams belonging to this endpoint
	MinMos *float64 `json:"minMos,omitempty"`


	// MinRFactor - The lowest R-factor value among all of the audio streams belonging to this endpoint
	MinRFactor *float64 `json:"minRFactor,omitempty"`


	// MaxLatencyMs - The maximum latency experienced by any audio stream belonging to this endpoint, in milliseconds
	MaxLatencyMs *int64 `json:"maxLatencyMs,omitempty"`


	// ReceivedPackets - The total number of packets received for all audio streams belonging to this endpoint (includes invalid, duplicate, and discarded packets)
	ReceivedPackets *int64 `json:"receivedPackets,omitempty"`


	// InvalidPackets - The total number of malformed or not RTP packets, unknown payload type, or discarded probation packets for all audio streams belonging to this endpoint
	InvalidPackets *int64 `json:"invalidPackets,omitempty"`


	// DiscardedPackets - The total number of packets received too late or too early, jitter queue overrun or underrun, for all audio streams belonging to this endpoint
	DiscardedPackets *int64 `json:"discardedPackets,omitempty"`


	// DuplicatePackets - The total number of packets received with the same sequence number as another one recently received (window of 64 packets), for all audio streams belonging to this endpoint
	DuplicatePackets *int64 `json:"duplicatePackets,omitempty"`


	// OverrunPackets - The total number of packets for which there was no room in the jitter queue when it was received, for all audio streams belonging to this endpoint (also counted in discarded)
	OverrunPackets *int64 `json:"overrunPackets,omitempty"`


	// UnderrunPackets - The total number of packets received after their timestamp/seqnum has been played out, for all audio streams belonging to this endpoint (also counted in discarded)
	UnderrunPackets *int64 `json:"underrunPackets,omitempty"`

}

// String returns a JSON representation of the model
func (o *Analyticsmediaendpointstat) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
