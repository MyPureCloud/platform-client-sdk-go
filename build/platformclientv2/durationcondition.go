package platformclientv2
import (
	"encoding/json"
)

// Durationcondition
type Durationcondition struct { 
	// DurationTarget
	DurationTarget *string `json:"durationTarget,omitempty"`


	// DurationOperator
	DurationOperator *string `json:"durationOperator,omitempty"`


	// DurationRange
	DurationRange *string `json:"durationRange,omitempty"`

}

// String returns a JSON representation of the model
func (o *Durationcondition) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
