package platformclientv2
import (
	"encoding/json"
)

// Items
type Items struct { 
	// VarType
	VarType *string `json:"type,omitempty"`


	// Pattern
	Pattern *string `json:"pattern,omitempty"`

}

// String returns a JSON representation of the model
func (o *Items) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
