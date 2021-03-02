package platformclientv2
import (
	"encoding/json"
)

// Context
type Context struct { 
	// Patterns - A list of one or more patterns to match.
	Patterns *[]Contextpattern `json:"patterns,omitempty"`

}

// String returns a JSON representation of the model
func (o *Context) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
