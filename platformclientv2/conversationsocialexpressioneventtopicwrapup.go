package platformclientv2
import (
	"time"
	"encoding/json"
	"strconv"
	"strings"
)

// Conversationsocialexpressioneventtopicwrapup
type Conversationsocialexpressioneventtopicwrapup struct { 
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
func (o *Conversationsocialexpressioneventtopicwrapup) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
