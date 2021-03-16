package platformclientv2
import (
	"time"
	"encoding/json"
)

// Workdaypointstrenditem
type Workdaypointstrenditem struct { 
	// DateWorkday - workday date for the points trend. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
	DateWorkday *time.Time `json:"dateWorkday,omitempty"`


	// Points - workday points for the date
	Points *float64 `json:"points,omitempty"`

}

// String returns a JSON representation of the model
func (o *Workdaypointstrenditem) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
