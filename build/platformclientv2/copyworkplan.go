package platformclientv2
import (
	"encoding/json"
)

// Copyworkplan - Information associated with a work plan thats created as a copy
type Copyworkplan struct { 
	// Name - Name of the copied work plan
	Name *string `json:"name,omitempty"`

}

// String returns a JSON representation of the model
func (o *Copyworkplan) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
