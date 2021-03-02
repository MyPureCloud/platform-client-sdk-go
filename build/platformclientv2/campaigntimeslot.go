package platformclientv2
import (
	"encoding/json"
)

// Campaigntimeslot
type Campaigntimeslot struct { 
	// StartTime - The start time of the interval as an ISO-8601 string, i.e. HH:mm:ss
	StartTime *string `json:"startTime,omitempty"`


	// StopTime - The end time of the interval as an ISO-8601 string, i.e. HH:mm:ss
	StopTime *string `json:"stopTime,omitempty"`


	// Day - The day of the interval. Valid values: [1-7], representing Monday through Sunday
	Day *int `json:"day,omitempty"`

}

// String returns a JSON representation of the model
func (o *Campaigntimeslot) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
