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

func (o *Journeypattern) MarshalJSON() ([]byte, error) {
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
		Criteria: o.Criteria,
		
		Count: o.Count,
		
		StreamType: o.StreamType,
		
		SessionType: o.SessionType,
		
		EventName: o.EventName,
		Alias:    (*Alias)(o),
	})
}

func (o *Journeypattern) UnmarshalJSON(b []byte) error {
	var JourneypatternMap map[string]interface{}
	err := json.Unmarshal(b, &JourneypatternMap)
	if err != nil {
		return err
	}
	
	if Criteria, ok := JourneypatternMap["criteria"].([]interface{}); ok {
		CriteriaString, _ := json.Marshal(Criteria)
		json.Unmarshal(CriteriaString, &o.Criteria)
	}
	
	if Count, ok := JourneypatternMap["count"].(float64); ok {
		CountInt := int(Count)
		o.Count = &CountInt
	}
	
	if StreamType, ok := JourneypatternMap["streamType"].(string); ok {
		o.StreamType = &StreamType
	}
	
	if SessionType, ok := JourneypatternMap["sessionType"].(string); ok {
		o.SessionType = &SessionType
	}
	
	if EventName, ok := JourneypatternMap["eventName"].(string); ok {
		o.EventName = &EventName
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Journeypattern) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
