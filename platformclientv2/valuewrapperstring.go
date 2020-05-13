package platformclientv2
import (
	"encoding/json"
)

// Valuewrapperstring - An object to provide context to nullable fields in PATCH requests
type Valuewrapperstring struct { 
	// Value - The value for the associated field
	Value *string `json:"value,omitempty"`

}

// String returns a JSON representation of the model
func (o *Valuewrapperstring) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
