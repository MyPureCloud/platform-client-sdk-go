package platformclientv2
import (
	"encoding/json"
)

// Timeslot
type Timeslot struct { 
	// StartTime - start time in xx:xx:xx.xxx format
	StartTime *string `json:"startTime,omitempty"`


	// StopTime - stop time in xx:xx:xx.xxx format
	StopTime *string `json:"stopTime,omitempty"`


	// Day - Day for this time slot, Monday = 1 ... Sunday = 7
	Day *int `json:"day,omitempty"`

}

// String returns a JSON representation of the model
func (o *Timeslot) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
