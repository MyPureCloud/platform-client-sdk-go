package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Changepasswordrequest
type Changepasswordrequest struct { 
	// NewPassword - The new password
	NewPassword *string `json:"newPassword,omitempty"`

}

// String returns a JSON representation of the model
func (o *Changepasswordrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
