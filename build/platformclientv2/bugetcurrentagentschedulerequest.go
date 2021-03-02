package platformclientv2
import (
	"time"
	"encoding/json"
)

// Bugetcurrentagentschedulerequest
type Bugetcurrentagentschedulerequest struct { 
	// StartDate - Start date of the range to search. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	StartDate *time.Time `json:"startDate,omitempty"`


	// EndDate - End date of the range to search. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	EndDate *time.Time `json:"endDate,omitempty"`

}

// String returns a JSON representation of the model
func (o *Bugetcurrentagentschedulerequest) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
