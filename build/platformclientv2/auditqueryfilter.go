package platformclientv2
import (
	"encoding/json"
)

// Auditqueryfilter
type Auditqueryfilter struct { 
	// Property - Name of the property to filter.
	Property *string `json:"property,omitempty"`


	// Value - Value of the property to filter.
	Value *string `json:"value,omitempty"`

}

// String returns a JSON representation of the model
func (o *Auditqueryfilter) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
