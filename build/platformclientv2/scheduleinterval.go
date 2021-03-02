package platformclientv2
import (
	"encoding/json"
)

// Scheduleinterval
type Scheduleinterval struct { 
	// Start - The scheduled start time as an ISO-8601 string, i.e yyyy-MM-ddTHH:mm:ss.SSSZ
	Start *string `json:"start,omitempty"`


	// End - The scheduled end time as an ISO-8601 string, i.e. yyyy-MM-ddTHH:mm:ss.SSSZ
	End *string `json:"end,omitempty"`

}

// String returns a JSON representation of the model
func (o *Scheduleinterval) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
