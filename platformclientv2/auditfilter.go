package platformclientv2
import (
	"encoding/json"
)

// Auditfilter
type Auditfilter struct { 
	// Name - The name of the field by which to filter.
	Name *string `json:"name,omitempty"`


	// VarType - The type of the filter, DATE or STRING.
	VarType *string `json:"type,omitempty"`


	// Operator - The operation that the filter performs.
	Operator *string `json:"operator,omitempty"`


	// Values - The values to make the filter comparison against.
	Values *[]string `json:"values,omitempty"`

}

// String returns a JSON representation of the model
func (o *Auditfilter) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
