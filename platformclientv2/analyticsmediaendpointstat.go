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

func (o *Analyticsmediaendpointstat) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Analyticsmediaendpointstat
	
	EventTime := new(string)
	if o.EventTime != nil {
		
		*EventTime = timeutil.Strftime(o.EventTime, "%Y-%m-%dT%H:%M:%S.%fZ")
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
		Codecs: o.Codecs,
		
		DiscardedPackets: o.DiscardedPackets,
		
		DuplicatePackets: o.DuplicatePackets,
		
		EventTime: EventTime,
		
		InvalidPackets: o.InvalidPackets,
		
		MaxLatencyMs: o.MaxLatencyMs,
		
		MinMos: o.MinMos,
		
		MinRFactor: o.MinRFactor,
		
		OverrunPackets: o.OverrunPackets,
		
		ReceivedPackets: o.ReceivedPackets,
		
		UnderrunPackets: o.UnderrunPackets,
		Alias:    (*Alias)(o),
	})
}

func (o *Analyticsmediaendpointstat) UnmarshalJSON(b []byte) error {
	var AnalyticsmediaendpointstatMap map[string]interface{}
	err := json.Unmarshal(b, &AnalyticsmediaendpointstatMap)
	if err != nil {
		return err
	}
	
	if Codecs, ok := AnalyticsmediaendpointstatMap["codecs"].([]interface{}); ok {
		CodecsString, _ := json.Marshal(Codecs)
		json.Unmarshal(CodecsString, &o.Codecs)
	}
	
	if DiscardedPackets, ok := AnalyticsmediaendpointstatMap["discardedPackets"].(float64); ok {
		DiscardedPacketsInt := int(DiscardedPackets)
		o.DiscardedPackets = &DiscardedPacketsInt
	}
	
	if DuplicatePackets, ok := AnalyticsmediaendpointstatMap["duplicatePackets"].(float64); ok {
		DuplicatePacketsInt := int(DuplicatePackets)
		o.DuplicatePackets = &DuplicatePacketsInt
	}
	
	if eventTimeString, ok := AnalyticsmediaendpointstatMap["eventTime"].(string); ok {
		EventTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", eventTimeString)
		o.EventTime = &EventTime
	}
	
	if InvalidPackets, ok := AnalyticsmediaendpointstatMap["invalidPackets"].(float64); ok {
		InvalidPacketsInt := int(InvalidPackets)
		o.InvalidPackets = &InvalidPacketsInt
	}
	
	if MaxLatencyMs, ok := AnalyticsmediaendpointstatMap["maxLatencyMs"].(float64); ok {
		MaxLatencyMsInt := int(MaxLatencyMs)
		o.MaxLatencyMs = &MaxLatencyMsInt
	}
	
	if MinMos, ok := AnalyticsmediaendpointstatMap["minMos"].(float64); ok {
		o.MinMos = &MinMos
	}
    
	if MinRFactor, ok := AnalyticsmediaendpointstatMap["minRFactor"].(float64); ok {
		o.MinRFactor = &MinRFactor
	}
    
	if OverrunPackets, ok := AnalyticsmediaendpointstatMap["overrunPackets"].(float64); ok {
		OverrunPacketsInt := int(OverrunPackets)
		o.OverrunPackets = &OverrunPacketsInt
	}
	
	if ReceivedPackets, ok := AnalyticsmediaendpointstatMap["receivedPackets"].(float64); ok {
		ReceivedPacketsInt := int(ReceivedPackets)
		o.ReceivedPackets = &ReceivedPacketsInt
	}
	
	if UnderrunPackets, ok := AnalyticsmediaendpointstatMap["underrunPackets"].(float64); ok {
		UnderrunPacketsInt := int(UnderrunPackets)
		o.UnderrunPackets = &UnderrunPacketsInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Analyticsmediaendpointstat) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
