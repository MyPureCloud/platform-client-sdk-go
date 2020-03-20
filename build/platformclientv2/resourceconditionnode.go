package platformclientv2
import (
	"encoding/json"
)

// Resourceconditionnode
type Resourceconditionnode struct { 
	// VariableName
	VariableName *string `json:"variableName,omitempty"`


	// Conjunction
	Conjunction *string `json:"conjunction,omitempty"`


	// Operator
	Operator *string `json:"operator,omitempty"`


	// Operands
	Operands *[]Resourceconditionvalue `json:"operands,omitempty"`


	// Terms
	Terms *[]Resourceconditionnode `json:"terms,omitempty"`

}

// String returns a JSON representation of the model
func (o *Resourceconditionnode) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
