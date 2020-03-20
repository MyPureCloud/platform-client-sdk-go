package platformclientv2
import (
	"encoding/json"
)

// Contentfilteritem
type Contentfilteritem struct { 
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
func (o *Contentfilteritem) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
