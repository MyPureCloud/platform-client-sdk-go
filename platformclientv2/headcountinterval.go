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

func (o *Headcountinterval) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Headcountinterval
	
	Interval := new(string)
	if o.Interval != nil {
		
		*Interval = timeutil.Strftime(o.Interval, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		Interval = nil
	}
	
	return json.Marshal(&struct { 
		Interval *string `json:"interval,omitempty"`
		
		Value *float64 `json:"value,omitempty"`
		*Alias
	}{ 
		Interval: Interval,
		
		Value: o.Value,
		Alias:    (*Alias)(o),
	})
}

func (o *Headcountinterval) UnmarshalJSON(b []byte) error {
	var HeadcountintervalMap map[string]interface{}
	err := json.Unmarshal(b, &HeadcountintervalMap)
	if err != nil {
		return err
	}
	
	if intervalString, ok := HeadcountintervalMap["interval"].(string); ok {
		Interval, _ := time.Parse("2006-01-02T15:04:05.999999Z", intervalString)
		o.Interval = &Interval
	}
	
	if Value, ok := HeadcountintervalMap["value"].(float64); ok {
		o.Value = &Value
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Headcountinterval) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
