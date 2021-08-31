package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Callabletime
type Callabletime struct { 
	// TimeSlots - The time intervals for which it is acceptable to place outbound calls.
	TimeSlots *[]Campaigntimeslot `json:"timeSlots,omitempty"`


	// TimeZoneId - The time zone for the time slots; for example, Africa/Abidjan
	TimeZoneId *string `json:"timeZoneId,omitempty"`

}

func (o *Callabletime) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Callabletime
	
	return json.Marshal(&struct { 
		TimeSlots *[]Campaigntimeslot `json:"timeSlots,omitempty"`
		
		TimeZoneId *string `json:"timeZoneId,omitempty"`
		*Alias
	}{ 
		TimeSlots: o.TimeSlots,
		
		TimeZoneId: o.TimeZoneId,
		Alias:    (*Alias)(o),
	})
}

func (o *Callabletime) UnmarshalJSON(b []byte) error {
	var CallabletimeMap map[string]interface{}
	err := json.Unmarshal(b, &CallabletimeMap)
	if err != nil {
		return err
	}
	
	if TimeSlots, ok := CallabletimeMap["timeSlots"].([]interface{}); ok {
		TimeSlotsString, _ := json.Marshal(TimeSlots)
		json.Unmarshal(TimeSlotsString, &o.TimeSlots)
	}
	
	if TimeZoneId, ok := CallabletimeMap["timeZoneId"].(string); ok {
		o.TimeZoneId = &TimeZoneId
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Callabletime) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
