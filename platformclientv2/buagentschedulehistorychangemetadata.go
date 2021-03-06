package platformclientv2
import (
	"time"
	"encoding/json"
)

// Buagentschedulehistorychangemetadata
type Buagentschedulehistorychangemetadata struct { 
	// DateModified - The timestamp of the schedule change. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateModified *time.Time `json:"dateModified,omitempty"`


	// ModifiedBy - The user that made the schedule change
	ModifiedBy *Userreference `json:"modifiedBy,omitempty"`

}

// String returns a JSON representation of the model
func (o *Buagentschedulehistorychangemetadata) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
