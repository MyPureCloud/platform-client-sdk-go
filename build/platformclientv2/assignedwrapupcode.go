package platformclientv2
import (
	"time"
	"encoding/json"
)

// Assignedwrapupcode
type Assignedwrapupcode struct { 
	// Code - The user configured wrap up code id.
	Code *string `json:"code,omitempty"`


	// Notes - Text entered by the agent to describe the call or disposition.
	Notes *string `json:"notes,omitempty"`


	// Tags - List of tags selected by the agent to describe the call or disposition.
	Tags *[]string `json:"tags,omitempty"`


	// DurationSeconds - The duration in seconds of the wrap-up segment.
	DurationSeconds *int32 `json:"durationSeconds,omitempty"`


	// EndTime - The timestamp when the wrap-up segment ended. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss.SSSZ
	EndTime *time.Time `json:"endTime,omitempty"`

}

// String returns a JSON representation of the model
func (o *Assignedwrapupcode) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
