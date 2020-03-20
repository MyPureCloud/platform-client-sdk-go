package platformclientv2
import (
	"encoding/json"
)

// Termattribute
type Termattribute struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// VarType
	VarType *string `json:"type,omitempty"`

}

// String returns a JSON representation of the model
func (o *Termattribute) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
