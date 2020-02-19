package platformclientv2
import (
	"encoding/json"
)

// Callforwardingeventtarget
type Callforwardingeventtarget struct { 
	// VarType
	VarType *string `json:"type,omitempty"`


	// Value
	Value *string `json:"value,omitempty"`

}

// String returns a JSON representation of the model
func (o *Callforwardingeventtarget) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
