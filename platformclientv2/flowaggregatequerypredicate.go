package platformclientv2
import (
	"encoding/json"
)

// Flowaggregatequerypredicate
type Flowaggregatequerypredicate struct { 
	// VarType - Optional type, can usually be inferred
	VarType *string `json:"type,omitempty"`


	// Dimension - Left hand side for dimension predicates
	Dimension *string `json:"dimension,omitempty"`


	// Operator - Optional operator, default is matches
	Operator *string `json:"operator,omitempty"`


	// Value - Right hand side for dimension predicates
	Value *string `json:"value,omitempty"`


	// VarRange - Right hand side for dimension predicates
	VarRange *Numericrange `json:"range,omitempty"`

}

// String returns a JSON representation of the model
func (o *Flowaggregatequerypredicate) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
