package platformclientv2
import (
	"encoding/json"
)

// Lexslot
type Lexslot struct { 
	// Name - The slot name
	Name *string `json:"name,omitempty"`


	// Description - The slot description
	Description *string `json:"description,omitempty"`


	// VarType - The slot type
	VarType *string `json:"type,omitempty"`


	// Priority - The priority of the slot
	Priority *int32 `json:"priority,omitempty"`

}

// String returns a JSON representation of the model
func (o *Lexslot) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
