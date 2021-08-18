package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Journeypattern
type Journeypattern struct { 
	// Criteria - A list of one or more criteria to satisfy.
	Criteria *[]Criteria `json:"criteria,omitempty"`


	// Count - The number of times the pattern must match.
	Count *int `json:"count,omitempty"`


	// StreamType - The stream type for which this pattern can be matched on.
	StreamType *string `json:"streamType,omitempty"`


	// SessionType - The session type for which this pattern can be matched on.
	SessionType *string `json:"sessionType,omitempty"`


	// EventName - The name of the event for which this pattern can be matched on.
	EventName *string `json:"eventName,omitempty"`

}

func (u *Journeypattern) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Journeypattern

	

	return json.Marshal(&struct { 
		Criteria *[]Criteria `json:"criteria,omitempty"`
		
		Count *int `json:"count,omitempty"`
		
		StreamType *string `json:"streamType,omitempty"`
		
		SessionType *string `json:"sessionType,omitempty"`
		
		EventName *string `json:"eventName,omitempty"`
		*Alias
	}{ 
		Criteria: u.Criteria,
		
		Count: u.Count,
		
		StreamType: u.StreamType,
		
		SessionType: u.SessionType,
		
		EventName: u.EventName,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Journeypattern) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
