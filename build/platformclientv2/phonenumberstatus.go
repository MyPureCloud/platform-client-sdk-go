package platformclientv2
import (
	"encoding/json"
)

// Phonenumberstatus
type Phonenumberstatus struct { 
	// Callable - Indicates whether or not a phone number is callable.
	Callable *bool `json:"callable,omitempty"`

}

// String returns a JSON representation of the model
func (o *Phonenumberstatus) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
