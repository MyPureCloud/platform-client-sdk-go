package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Workdayvaluestrenditem
type Workdayvaluestrenditem struct { 
	// DateWorkday - The workday for the metric value. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
	DateWorkday *time.Time `json:"dateWorkday,omitempty"`


	// Value - The metric value
	Value *float64 `json:"value,omitempty"`

}

func (o *Workdayvaluestrenditem) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Workdayvaluestrenditem
	
	DateWorkday := new(string)
	if o.DateWorkday != nil {
		*DateWorkday = timeutil.Strftime(o.DateWorkday, "%Y-%m-%d")
	} else {
		DateWorkday = nil
	}
	
	return json.Marshal(&struct { 
		DateWorkday *string `json:"dateWorkday,omitempty"`
		
		Value *float64 `json:"value,omitempty"`
		*Alias
	}{ 
		DateWorkday: DateWorkday,
		
		Value: o.Value,
		Alias:    (*Alias)(o),
	})
}

func (o *Workdayvaluestrenditem) UnmarshalJSON(b []byte) error {
	var WorkdayvaluestrenditemMap map[string]interface{}
	err := json.Unmarshal(b, &WorkdayvaluestrenditemMap)
	if err != nil {
		return err
	}
	
	if dateWorkdayString, ok := WorkdayvaluestrenditemMap["dateWorkday"].(string); ok {
		DateWorkday, _ := time.Parse("2006-01-02", dateWorkdayString)
		o.DateWorkday = &DateWorkday
	}
	
	if Value, ok := WorkdayvaluestrenditemMap["value"].(float64); ok {
		o.Value = &Value
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Workdayvaluestrenditem) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
