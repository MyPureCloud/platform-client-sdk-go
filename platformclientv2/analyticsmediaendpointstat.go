package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Analyticsmediaendpointstat
type Analyticsmediaendpointstat struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
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

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Analyticsmediaendpointstat) SetField(field string, fieldValue interface{}) {
	// Get Value object for field
	target := reflect.ValueOf(o)
	targetField := reflect.Indirect(target).FieldByName(field)

	// Set value
	if fieldValue != nil {
		targetField.Set(reflect.ValueOf(fieldValue))
	} else {
		// Must create a new Value (creates **type) then get its element (*type), which will be nil pointer of the appropriate type
		x := reflect.Indirect(reflect.New(targetField.Type()))
		targetField.Set(x)
	}

	// Add field to set field names list
	if o.SetFieldNames == nil {
		o.SetFieldNames = make(map[string]bool)
	}
	o.SetFieldNames[field] = true
}

func (o Analyticsmediaendpointstat) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "EventTime", }
		localDateTimeFields := []string{  }
		dateFields := []string{  }

		// Construct object
		newObj := make(map[string]interface{})
		for fieldName := range o.SetFieldNames {
			// Get initial field value
			fieldValue := val.FieldByName(fieldName).Interface()

			// Apply value formatting overrides
			if fieldValue == nil || reflect.ValueOf(fieldValue).IsNil()  {
				// Do nothing. Just catching this case to avoid trying to custom serialize a nil value
			} else if contains(dateTimeFields, fieldName) {
				fieldValue = timeutil.Strftime(toTime(fieldValue), "%Y-%m-%dT%H:%M:%S.%fZ")
			} else if contains(localDateTimeFields, fieldName) {
				fieldValue = timeutil.Strftime(toTime(fieldValue), "%Y-%m-%dT%H:%M:%S.%f")
			} else if contains(dateFields, fieldName) {
				fieldValue = timeutil.Strftime(toTime(fieldValue), "%Y-%m-%d")
			}

			// Assign value to field using JSON tag name
			newObj[getFieldName(reflect.TypeOf(&o), fieldName)] = fieldValue
		}

		// Marshal and return dynamically constructed interface
		return json.Marshal(newObj)
	}

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
		Alias
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
		Alias:    (Alias)(o),
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
