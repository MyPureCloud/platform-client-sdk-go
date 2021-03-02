package platformclientv2
import (
	"encoding/json"
)

// Keyvalue
type Keyvalue struct { 
	// Key - Key for free-form data.
	Key *string `json:"key,omitempty"`


	// Value - Value for free-form data.
	Value *string `json:"value,omitempty"`

}

// String returns a JSON representation of the model
func (o *Keyvalue) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
