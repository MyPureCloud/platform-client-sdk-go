package platformclientv2
import (
	"encoding/json"
)

// Schedulegenerationresultsummary
type Schedulegenerationresultsummary struct { 
	// Failed - Whether the schedule generation run failed
	Failed *bool `json:"failed,omitempty"`


	// RunId - The run ID for the schedule generation. Reference this when requesting support
	RunId *string `json:"runId,omitempty"`


	// MessageCount - The number of schedule generation messages for this schedule generation run
	MessageCount *int `json:"messageCount,omitempty"`

}

// String returns a JSON representation of the model
func (o *Schedulegenerationresultsummary) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
