package platformclientv2
import (
	"encoding/json"
)

// Namedentitytypedefinition
type Namedentitytypedefinition struct { 
	// Name - The name of the entity type.
	Name *string `json:"name,omitempty"`


	// Description - Description of the of the named entity type.
	Description *string `json:"description,omitempty"`


	// Mechanism - The mechanism enabling detection of the named entity type.
	Mechanism *Namedentitytypemechanism `json:"mechanism,omitempty"`

}

// String returns a JSON representation of the model
func (o *Namedentitytypedefinition) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
