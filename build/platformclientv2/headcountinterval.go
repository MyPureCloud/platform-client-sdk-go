package platformclientv2
import (
	"time"
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

// String returns a JSON representation of the model
func (o *Headcountinterval) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
