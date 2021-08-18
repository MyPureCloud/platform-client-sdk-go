package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Changemypasswordrequest
type Changemypasswordrequest struct { 
	// NewPassword - The new password
	NewPassword *string `json:"newPassword,omitempty"`


	// OldPassword - Your current password
	OldPassword *string `json:"oldPassword,omitempty"`

}

func (u *Changemypasswordrequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Changemypasswordrequest

	

	return json.Marshal(&struct { 
		NewPassword *string `json:"newPassword,omitempty"`
		
		OldPassword *string `json:"oldPassword,omitempty"`
		*Alias
	}{ 
		NewPassword: u.NewPassword,
		
		OldPassword: u.OldPassword,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Changemypasswordrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
