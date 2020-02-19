package platformclientv2
import (
	"encoding/json"
)

// Wfmbuscheduleruntopicbuschedulereference
type Wfmbuscheduleruntopicbuschedulereference struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// WeekDate
	WeekDate *Wfmbuscheduleruntopiclocaldate `json:"weekDate,omitempty"`

}

// String returns a JSON representation of the model
func (o *Wfmbuscheduleruntopicbuschedulereference) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
