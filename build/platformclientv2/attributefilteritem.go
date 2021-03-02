package platformclientv2
import (
	"encoding/json"
)

// Attributefilteritem
type Attributefilteritem struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// Operator
	Operator *string `json:"operator,omitempty"`


	// Values
	Values *[]string `json:"values,omitempty"`

}

// String returns a JSON representation of the model
func (o *Attributefilteritem) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
