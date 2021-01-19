package platformclientv2
import (
	"encoding/json"
)

// Journeypattern
type Journeypattern struct { 
	// Criteria - A list of one or more criteria to satisfy.
	Criteria *[]Criteria `json:"criteria,omitempty"`


	// Count - The number of times the pattern must match.
	Count *int32 `json:"count,omitempty"`


	// StreamType - The stream type for which this pattern can be matched on.
	StreamType *string `json:"streamType,omitempty"`


	// SessionType - The session type for which this pattern can be matched on.
	SessionType *string `json:"sessionType,omitempty"`


	// EventName - The name of the event for which this pattern can be matched on.
	EventName *string `json:"eventName,omitempty"`

}

// String returns a JSON representation of the model
func (o *Journeypattern) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
