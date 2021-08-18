package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Scimemail - Defines a SCIM email address.
type Scimemail struct { 
	// Value - The email address. Is immutable if \"type\" is set to \"other\".
	Value *string `json:"value,omitempty"`


	// VarType - The type of email address. \"value\" is immutable if \"type\" is set to \"other\".
	VarType *string `json:"type,omitempty"`


	// Primary - Indicates whether the email address is the primary email address.
	Primary *bool `json:"primary,omitempty"`

}

func (u *Scimemail) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Scimemail

	

	return json.Marshal(&struct { 
		Value *string `json:"value,omitempty"`
		
		VarType *string `json:"type,omitempty"`
		
		Primary *bool `json:"primary,omitempty"`
		*Alias
	}{ 
		Value: u.Value,
		
		VarType: u.VarType,
		
		Primary: u.Primary,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Scimemail) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
