package platformclientv2
import (
	"encoding/json"
)

// Copyworkplanrotationrequest
type Copyworkplanrotationrequest struct { 
	// Name - Name to apply to the new copy of the work plan rotation
	Name *string `json:"name,omitempty"`

}

// String returns a JSON representation of the model
func (o *Copyworkplanrotationrequest) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
