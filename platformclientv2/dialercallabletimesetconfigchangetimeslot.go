package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Dialercallabletimesetconfigchangetimeslot
type Dialercallabletimesetconfigchangetimeslot struct { 
	// StartTime - The start time of this time slot
	StartTime *string `json:"startTime,omitempty"`


	// StopTime - The stop time of this time slot
	StopTime *string `json:"stopTime,omitempty"`


	// Day - The day this time slot applies
	Day *int `json:"day,omitempty"`

}

func (o *Dialercallabletimesetconfigchangetimeslot) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Dialercallabletimesetconfigchangetimeslot
	
	return json.Marshal(&struct { 
		StartTime *string `json:"startTime,omitempty"`
		
		StopTime *string `json:"stopTime,omitempty"`
		
		Day *int `json:"day,omitempty"`
		*Alias
	}{ 
		StartTime: o.StartTime,
		
		StopTime: o.StopTime,
		
		Day: o.Day,
		Alias:    (*Alias)(o),
	})
}

func (o *Dialercallabletimesetconfigchangetimeslot) UnmarshalJSON(b []byte) error {
	var DialercallabletimesetconfigchangetimeslotMap map[string]interface{}
	err := json.Unmarshal(b, &DialercallabletimesetconfigchangetimeslotMap)
	if err != nil {
		return err
	}
	
	if StartTime, ok := DialercallabletimesetconfigchangetimeslotMap["startTime"].(string); ok {
		o.StartTime = &StartTime
	}
    
	if StopTime, ok := DialercallabletimesetconfigchangetimeslotMap["stopTime"].(string); ok {
		o.StopTime = &StopTime
	}
    
	if Day, ok := DialercallabletimesetconfigchangetimeslotMap["day"].(float64); ok {
		DayInt := int(Day)
		o.Day = &DayInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Dialercallabletimesetconfigchangetimeslot) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
