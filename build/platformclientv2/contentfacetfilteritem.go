package platformclientv2
import (
	"encoding/json"
)

// Contentfacetfilteritem
type Contentfacetfilteritem struct { 
	// Name
	Name *string `json:"name,omitempty"`


	// VarType
	VarType *string `json:"type,omitempty"`


	// Operator
	Operator *string `json:"operator,omitempty"`


	// Values
	Values *[]string `json:"values,omitempty"`

}

// String returns a JSON representation of the model
func (o *Contentfacetfilteritem) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
