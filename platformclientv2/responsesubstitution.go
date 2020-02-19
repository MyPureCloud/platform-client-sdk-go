package platformclientv2
import (
	"encoding/json"
)

// Responsesubstitution - Contains information about the substitutions associated with a response.
type Responsesubstitution struct { 
	// Id - Response substitution identifier.
	Id *string `json:"id,omitempty"`


	// Description - Response substitution description.
	Description *string `json:"description,omitempty"`


	// DefaultValue - Response substitution default value.
	DefaultValue *string `json:"defaultValue,omitempty"`

}

// String returns a JSON representation of the model
func (o *Responsesubstitution) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
