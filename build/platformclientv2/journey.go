package platformclientv2
import (
	"encoding/json"
)

// Journey
type Journey struct { 
	// Patterns - A list of zero or more patterns to match.
	Patterns *[]Journeypattern `json:"patterns,omitempty"`

}

// String returns a JSON representation of the model
func (o *Journey) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
