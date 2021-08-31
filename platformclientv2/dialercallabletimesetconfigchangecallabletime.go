package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Dialercallabletimesetconfigchangecallabletime
type Dialercallabletimesetconfigchangecallabletime struct { 
	// TimeSlots
	TimeSlots *[]Dialercallabletimesetconfigchangetimeslot `json:"timeSlots,omitempty"`


	// TimeZoneId
	TimeZoneId *string `json:"timeZoneId,omitempty"`


	// AdditionalProperties
	AdditionalProperties *interface{} `json:"additionalProperties,omitempty"`

}

func (o *Dialercallabletimesetconfigchangecallabletime) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Dialercallabletimesetconfigchangecallabletime
	
	return json.Marshal(&struct { 
		TimeSlots *[]Dialercallabletimesetconfigchangetimeslot `json:"timeSlots,omitempty"`
		
		TimeZoneId *string `json:"timeZoneId,omitempty"`
		
		AdditionalProperties *interface{} `json:"additionalProperties,omitempty"`
		*Alias
	}{ 
		TimeSlots: o.TimeSlots,
		
		TimeZoneId: o.TimeZoneId,
		
		AdditionalProperties: o.AdditionalProperties,
		Alias:    (*Alias)(o),
	})
}

func (o *Dialercallabletimesetconfigchangecallabletime) UnmarshalJSON(b []byte) error {
	var DialercallabletimesetconfigchangecallabletimeMap map[string]interface{}
	err := json.Unmarshal(b, &DialercallabletimesetconfigchangecallabletimeMap)
	if err != nil {
		return err
	}
	
	if TimeSlots, ok := DialercallabletimesetconfigchangecallabletimeMap["timeSlots"].([]interface{}); ok {
		TimeSlotsString, _ := json.Marshal(TimeSlots)
		json.Unmarshal(TimeSlotsString, &o.TimeSlots)
	}
	
	if TimeZoneId, ok := DialercallabletimesetconfigchangecallabletimeMap["timeZoneId"].(string); ok {
		o.TimeZoneId = &TimeZoneId
	}
	
	if AdditionalProperties, ok := DialercallabletimesetconfigchangecallabletimeMap["additionalProperties"].(map[string]interface{}); ok {
		AdditionalPropertiesString, _ := json.Marshal(AdditionalProperties)
		json.Unmarshal(AdditionalPropertiesString, &o.AdditionalProperties)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Dialercallabletimesetconfigchangecallabletime) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
