package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Timeallowed
type Timeallowed struct { 
	// TimeSlots
	TimeSlots *[]Timeslot `json:"timeSlots,omitempty"`


	// TimeZoneId
	TimeZoneId *string `json:"timeZoneId,omitempty"`


	// Empty
	Empty *bool `json:"empty,omitempty"`

}

func (o *Timeallowed) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Timeallowed
	
	return json.Marshal(&struct { 
		TimeSlots *[]Timeslot `json:"timeSlots,omitempty"`
		
		TimeZoneId *string `json:"timeZoneId,omitempty"`
		
		Empty *bool `json:"empty,omitempty"`
		*Alias
	}{ 
		TimeSlots: o.TimeSlots,
		
		TimeZoneId: o.TimeZoneId,
		
		Empty: o.Empty,
		Alias:    (*Alias)(o),
	})
}

func (o *Timeallowed) UnmarshalJSON(b []byte) error {
	var TimeallowedMap map[string]interface{}
	err := json.Unmarshal(b, &TimeallowedMap)
	if err != nil {
		return err
	}
	
	if TimeSlots, ok := TimeallowedMap["timeSlots"].([]interface{}); ok {
		TimeSlotsString, _ := json.Marshal(TimeSlots)
		json.Unmarshal(TimeSlotsString, &o.TimeSlots)
	}
	
	if TimeZoneId, ok := TimeallowedMap["timeZoneId"].(string); ok {
		o.TimeZoneId = &TimeZoneId
	}
    
	if Empty, ok := TimeallowedMap["empty"].(bool); ok {
		o.Empty = &Empty
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Timeallowed) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
