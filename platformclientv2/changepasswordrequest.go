package platformclientv2
import (
	"encoding/json"
)

// Changepasswordrequest
type Changepasswordrequest struct { 
	// NewPassword - The new password
	NewPassword *string `json:"newPassword,omitempty"`

}

// String returns a JSON representation of the model
func (o *Changepasswordrequest) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
