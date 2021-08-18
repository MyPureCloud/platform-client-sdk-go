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

func (u *Dialercallabletimesetconfigchangetimeslot) MarshalJSON() ([]byte, error) {
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
		StartTime: u.StartTime,
		
		StopTime: u.StopTime,
		
		Day: u.Day,
		
		AdditionalProperties: u.AdditionalProperties,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Dialercallabletimesetconfigchangetimeslot) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
