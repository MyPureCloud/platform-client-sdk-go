package platformclientv2
import (
	"encoding/json"
)

// Templateparameter
type Templateparameter struct { 
	// Id - Response substitution identifier
	Id *string `json:"id,omitempty"`


	// Value - Response substitution value
	Value *string `json:"value,omitempty"`

}

// String returns a JSON representation of the model
func (o *Templateparameter) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
