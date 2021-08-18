package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Changepasswordrequest
type Changepasswordrequest struct { 
	// NewPassword - The new password
	NewPassword *string `json:"newPassword,omitempty"`

}

func (u *Changepasswordrequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Changepasswordrequest

	

	return json.Marshal(&struct { 
		NewPassword *string `json:"newPassword,omitempty"`
		*Alias
	}{ 
		NewPassword: u.NewPassword,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Changepasswordrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
