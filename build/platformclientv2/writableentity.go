package platformclientv2
import (
	"encoding/json"
)

// Writableentity
type Writableentity struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`

}

// String returns a JSON representation of the model
func (o *Writableentity) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
