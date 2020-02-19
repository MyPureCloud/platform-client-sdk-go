package platformclientv2
import (
	"encoding/json"
)

// Userparam
type Userparam struct { 
	// Key
	Key *string `json:"key,omitempty"`


	// Value
	Value *string `json:"value,omitempty"`

}

// String returns a JSON representation of the model
func (o *Userparam) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
