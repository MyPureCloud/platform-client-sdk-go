package platformclientv2
import (
	"encoding/json"
)

// Socialhandle
type Socialhandle struct { 
	// VarType
	VarType *string `json:"type,omitempty"`


	// Value
	Value *string `json:"value,omitempty"`

}

// String returns a JSON representation of the model
func (o *Socialhandle) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
