package platformclientv2
import (
	"time"
	"encoding/json"
)

// Wrapup
type Wrapup struct { 
	// Code - The user configured wrap up code id.
	Code *string `json:"code,omitempty"`


	// Name - The user configured wrap up code name.
	Name *string `json:"name,omitempty"`


	// Notes - Text entered by the agent to describe the call or disposition.
	Notes *string `json:"notes,omitempty"`


	// Tags - List of tags selected by the agent to describe the call or disposition.
	Tags *[]string `json:"tags,omitempty"`


	// DurationSeconds - The length of time in seconds that the agent spent doing after call work.
	DurationSeconds *int32 `json:"durationSeconds,omitempty"`


	// EndTime - The timestamp when the wrapup was finished. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	EndTime *time.Time `json:"endTime,omitempty"`


	// Provisional - Indicates if this is a pending save and should not require a code to be specified.  This allows someone to save some temporary wrapup that will be used later.
	Provisional *bool `json:"provisional,omitempty"`

}

// String returns a JSON representation of the model
func (o *Wrapup) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
