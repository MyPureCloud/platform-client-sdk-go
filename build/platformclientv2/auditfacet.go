package platformclientv2
import (
	"encoding/json"
)

// Auditfacet
type Auditfacet struct { 
	// Name - The name of the field on which to facet.
	Name *string `json:"name,omitempty"`


	// VarType - The type of the facet, DATE or STRING.
	VarType *string `json:"type,omitempty"`

}

// String returns a JSON representation of the model
func (o *Auditfacet) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
