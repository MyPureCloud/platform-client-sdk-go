package platformclientv2
import (
	"encoding/json"
)

// Namedentitytypemechanism
type Namedentitytypemechanism struct { 
	// Items - The items that define the named entity type.
	Items *[]Namedentitytypeitem `json:"items,omitempty"`


	// Restricted - Whether the named entity type is restricted to the items provided. Default: false
	Restricted *bool `json:"restricted,omitempty"`


	// VarType - The type of the mechanism.
	VarType *string `json:"type,omitempty"`

}

// String returns a JSON representation of the model
func (o *Namedentitytypemechanism) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
