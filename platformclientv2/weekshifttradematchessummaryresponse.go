package platformclientv2
import (
	"time"
	"encoding/json"
)

// Weekshifttradematchessummaryresponse
type Weekshifttradematchessummaryresponse struct { 
	// WeekDate - The schedule week date in yyyy-MM-dd format. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
	WeekDate *time.Time `json:"weekDate,omitempty"`


	// Count - The number of trades in the Matched state for the given week
	Count *int `json:"count,omitempty"`

}

// String returns a JSON representation of the model
func (o *Weekshifttradematchessummaryresponse) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
