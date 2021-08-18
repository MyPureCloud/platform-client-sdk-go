package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Headcountinterval - Headcount interval information for schedule
type Headcountinterval struct { 
	// Interval - The start date-time for this headcount interval in ISO-8601 format, must be within the 8 day schedule
	Interval *time.Time `json:"interval,omitempty"`


	// Value - Headcount value for this interval
	Value *float64 `json:"value,omitempty"`

}

func (u *Headcountinterval) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Headcountinterval

	
	Interval := new(string)
	if u.Interval != nil {
		
		*Interval = timeutil.Strftime(u.Interval, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		Interval = nil
	}
	

	return json.Marshal(&struct { 
		Interval *string `json:"interval,omitempty"`
		
		Value *float64 `json:"value,omitempty"`
		*Alias
	}{ 
		Interval: Interval,
		
		Value: u.Value,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Headcountinterval) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
