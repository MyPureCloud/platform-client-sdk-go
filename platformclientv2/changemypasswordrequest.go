package platformclientv2
import (
	"encoding/json"
)

// Changemypasswordrequest
type Changemypasswordrequest struct { 
	// NewPassword - The new password
	NewPassword *string `json:"newPassword,omitempty"`


	// OldPassword - Your current password
	OldPassword *string `json:"oldPassword,omitempty"`

}

// String returns a JSON representation of the model
func (o *Changemypasswordrequest) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
