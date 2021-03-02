package platformclientv2
import (
	"encoding/json"
)

// Resolutiondetailquerypredicate
type Resolutiondetailquerypredicate struct { 
	// VarType - Optional type, can usually be inferred
	VarType *string `json:"type,omitempty"`


	// Metric - Left hand side for metric predicates
	Metric *string `json:"metric,omitempty"`


	// Operator - Optional operator, default is matches
	Operator *string `json:"operator,omitempty"`


	// Value - Right hand side for metric predicates
	Value *string `json:"value,omitempty"`


	// VarRange - Right hand side for metric predicates
	VarRange *Numericrange `json:"range,omitempty"`

}

// String returns a JSON representation of the model
func (o *Resolutiondetailquerypredicate) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
