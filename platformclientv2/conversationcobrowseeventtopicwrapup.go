package platformclientv2
import (
	"time"
	"encoding/json"
)

// Conversationcobrowseeventtopicwrapup
type Conversationcobrowseeventtopicwrapup struct { 
	// Code
	Code *string `json:"code,omitempty"`


	// Notes
	Notes *string `json:"notes,omitempty"`


	// Tags
	Tags *[]string `json:"tags,omitempty"`


	// DurationSeconds
	DurationSeconds *int `json:"durationSeconds,omitempty"`


	// EndTime
	EndTime *time.Time `json:"endTime,omitempty"`


	// AdditionalProperties
	AdditionalProperties *map[string]interface{} `json:"additionalProperties,omitempty"`

}

// String returns a JSON representation of the model
func (o *Conversationcobrowseeventtopicwrapup) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
