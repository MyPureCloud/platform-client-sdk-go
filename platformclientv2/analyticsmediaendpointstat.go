package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Analyticsmediaendpointstat
type Analyticsmediaendpointstat struct { 
	// Codecs - The MIME type(s) of the audio encodings used by the audio streams belonging to this endpoint
	Codecs *[]string `json:"codecs,omitempty"`


	// DiscardedPackets - The total number of packets received too late or too early, jitter queue overrun or underrun, for all audio streams belonging to this endpoint
	DiscardedPackets *int `json:"discardedPackets,omitempty"`


	// DuplicatePackets - The total number of packets received with the same sequence number as another one recently received (window of 64 packets), for all audio streams belonging to this endpoint
	DuplicatePackets *int `json:"duplicatePackets,omitempty"`


	// EventTime - Specifies when an event occurred. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	EventTime *time.Time `json:"eventTime,omitempty"`


	// InvalidPackets - The total number of malformed or not RTP packets, unknown payload type, or discarded probation packets for all audio streams belonging to this endpoint
	InvalidPackets *int `json:"invalidPackets,omitempty"`


	// MaxLatencyMs - The maximum latency experienced by any audio stream belonging to this endpoint, in milliseconds
	MaxLatencyMs *int `json:"maxLatencyMs,omitempty"`


	// MinMos - The lowest estimated average MOS among all the audio streams belonging to this endpoint
	MinMos *float64 `json:"minMos,omitempty"`


	// MinRFactor - The lowest R-factor value among all of the audio streams belonging to this endpoint
	MinRFactor *float64 `json:"minRFactor,omitempty"`


	// OverrunPackets - The total number of packets for which there was no room in the jitter queue when it was received, for all audio streams belonging to this endpoint (also counted in discarded)
	OverrunPackets *int `json:"overrunPackets,omitempty"`


	// ReceivedPackets - The total number of packets received for all audio streams belonging to this endpoint (includes invalid, duplicate, and discarded packets)
	ReceivedPackets *int `json:"receivedPackets,omitempty"`


	// UnderrunPackets - The total number of packets received after their timestamp/seqnum has been played out, for all audio streams belonging to this endpoint (also counted in discarded)
	UnderrunPackets *int `json:"underrunPackets,omitempty"`

}

func (u *Analyticsmediaendpointstat) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Analyticsmediaendpointstat

	
	EventTime := new(string)
	if u.EventTime != nil {
		
		*EventTime = timeutil.Strftime(u.EventTime, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		EventTime = nil
	}
	

	return json.Marshal(&struct { 
		Codecs *[]string `json:"codecs,omitempty"`
		
		DiscardedPackets *int `json:"discardedPackets,omitempty"`
		
		DuplicatePackets *int `json:"duplicatePackets,omitempty"`
		
		EventTime *string `json:"eventTime,omitempty"`
		
		InvalidPackets *int `json:"invalidPackets,omitempty"`
		
		MaxLatencyMs *int `json:"maxLatencyMs,omitempty"`
		
		MinMos *float64 `json:"minMos,omitempty"`
		
		MinRFactor *float64 `json:"minRFactor,omitempty"`
		
		OverrunPackets *int `json:"overrunPackets,omitempty"`
		
		ReceivedPackets *int `json:"receivedPackets,omitempty"`
		
		UnderrunPackets *int `json:"underrunPackets,omitempty"`
		*Alias
	}{ 
		Codecs: u.Codecs,
		
		DiscardedPackets: u.DiscardedPackets,
		
		DuplicatePackets: u.DuplicatePackets,
		
		EventTime: EventTime,
		
		InvalidPackets: u.InvalidPackets,
		
		MaxLatencyMs: u.MaxLatencyMs,
		
		MinMos: u.MinMos,
		
		MinRFactor: u.MinRFactor,
		
		OverrunPackets: u.OverrunPackets,
		
		ReceivedPackets: u.ReceivedPackets,
		
		UnderrunPackets: u.UnderrunPackets,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Analyticsmediaendpointstat) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
