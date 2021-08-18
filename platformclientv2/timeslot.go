package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Timeslot
type Timeslot struct { 
	// StartTime - start time in xx:xx:xx.xxx format
	StartTime *string `json:"startTime,omitempty"`


	// StopTime - stop time in xx:xx:xx.xxx format
	StopTime *string `json:"stopTime,omitempty"`


	// Day - Day for this time slot, Monday = 1 ... Sunday = 7
	Day *int `json:"day,omitempty"`

}

func (u *Timeslot) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Timeslot

	

	return json.Marshal(&struct { 
		StartTime *string `json:"startTime,omitempty"`
		
		StopTime *string `json:"stopTime,omitempty"`
		
		Day *int `json:"day,omitempty"`
		*Alias
	}{ 
		StartTime: u.StartTime,
		
		StopTime: u.StopTime,
		
		Day: u.Day,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Timeslot) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
