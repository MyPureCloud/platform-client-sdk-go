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

func (o *Changemypasswordrequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Changemypasswordrequest
	
	return json.Marshal(&struct { 
		NewPassword *string `json:"newPassword,omitempty"`
		
		OldPassword *string `json:"oldPassword,omitempty"`
		*Alias
	}{ 
		NewPassword: o.NewPassword,
		
		OldPassword: o.OldPassword,
		Alias:    (*Alias)(o),
	})
}

func (o *Changemypasswordrequest) UnmarshalJSON(b []byte) error {
	var ChangemypasswordrequestMap map[string]interface{}
	err := json.Unmarshal(b, &ChangemypasswordrequestMap)
	if err != nil {
		return err
	}
	
	if NewPassword, ok := ChangemypasswordrequestMap["newPassword"].(string); ok {
		o.NewPassword = &NewPassword
	}
    
	if OldPassword, ok := ChangemypasswordrequestMap["oldPassword"].(string); ok {
		o.OldPassword = &OldPassword
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Changemypasswordrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
