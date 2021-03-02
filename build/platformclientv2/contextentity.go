package platformclientv2
import (
	"encoding/json"
)

// Contextentity
type Contextentity struct { 
	// Name - The name of the entity.
	Name *string `json:"name,omitempty"`

}

// String returns a JSON representation of the model
func (o *Contextentity) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
