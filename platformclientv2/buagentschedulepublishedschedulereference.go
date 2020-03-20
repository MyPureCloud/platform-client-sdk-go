package platformclientv2
import (
	"time"
	"encoding/json"
)

// Buagentschedulepublishedschedulereference
type Buagentschedulepublishedschedulereference struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// WeekDate - The start week date for this schedule. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
	WeekDate *time.Time `json:"weekDate,omitempty"`


	// WeekCount - The number of weeks encompassed by the schedule
	WeekCount *int32 `json:"weekCount,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

// String returns a JSON representation of the model
func (o *Buagentschedulepublishedschedulereference) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
