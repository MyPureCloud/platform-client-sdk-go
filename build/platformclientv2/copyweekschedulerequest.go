package platformclientv2
import (
	"encoding/json"
)

// Copyweekschedulerequest
type Copyweekschedulerequest struct { 
	// Description - Description of the copied week schedule
	Description *string `json:"description,omitempty"`


	// WeekDate - Week in yyyy-MM-dd format to which the schedule is copied
	WeekDate *string `json:"weekDate,omitempty"`

}

// String returns a JSON representation of the model
func (o *Copyweekschedulerequest) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
