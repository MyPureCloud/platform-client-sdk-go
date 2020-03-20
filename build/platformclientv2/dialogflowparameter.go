package platformclientv2
import (
	"encoding/json"
)

// Dialogflowparameter
type Dialogflowparameter struct { 
	// Name - The parameter name
	Name *string `json:"name,omitempty"`


	// VarType - The parameter type
	VarType *string `json:"type,omitempty"`

}

// String returns a JSON representation of the model
func (o *Dialogflowparameter) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
