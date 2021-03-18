package platformclientv2
import (
	"time"
	"encoding/json"
	"strconv"
	"strings"
)

// Bucopyschedulerequest
type Bucopyschedulerequest struct { 
	// Description - The description for the new schedule
	Description *string `json:"description,omitempty"`


	// WeekDate - The start weekDate for the new copy of the schedule. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
	WeekDate *time.Time `json:"weekDate,omitempty"`

}

// String returns a JSON representation of the model
func (o *Bucopyschedulerequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
