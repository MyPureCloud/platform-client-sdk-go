package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Dialercallabletimesetconfigchangetimeslot
type Dialercallabletimesetconfigchangetimeslot struct { 
	// StartTime
	StartTime *string `json:"startTime,omitempty"`


	// StopTime
	StopTime *string `json:"stopTime,omitempty"`


	// Day
	Day *int `json:"day,omitempty"`


	// AdditionalProperties
	AdditionalProperties *interface{} `json:"additionalProperties,omitempty"`

}

func (o *Dialercallabletimesetconfigchangetimeslot) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Dialercallabletimesetconfigchangetimeslot
	
	return json.Marshal(&struct { 
		StartTime *string `json:"startTime,omitempty"`
		
		StopTime *string `json:"stopTime,omitempty"`
		
		Day *int `json:"day,omitempty"`
		
		AdditionalProperties *interface{} `json:"additionalProperties,omitempty"`
		*Alias
	}{ 
		StartTime: o.StartTime,
		
		StopTime: o.StopTime,
		
		Day: o.Day,
		
		AdditionalProperties: o.AdditionalProperties,
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
	
	if AdditionalProperties, ok := DialercallabletimesetconfigchangetimeslotMap["additionalProperties"].(map[string]interface{}); ok {
		AdditionalPropertiesString, _ := json.Marshal(AdditionalProperties)
		json.Unmarshal(AdditionalPropertiesString, &o.AdditionalProperties)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Dialercallabletimesetconfigchangetimeslot) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
