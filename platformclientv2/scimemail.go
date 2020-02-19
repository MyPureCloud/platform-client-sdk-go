package platformclientv2
import (
	"encoding/json"
)

// Scimemail - Defines a SCIM email address.
type Scimemail struct { 
	// Value - The email address.
	Value *string `json:"value,omitempty"`


	// VarType - The type of email address.
	VarType *string `json:"type,omitempty"`


	// Primary - Indicates whether the email address is the primary email address.
	Primary *bool `json:"primary,omitempty"`

}

// String returns a JSON representation of the model
func (o *Scimemail) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
