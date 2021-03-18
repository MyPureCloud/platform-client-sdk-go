package platformclientv2
import (
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

// String returns a JSON representation of the model
func (o *Scimemail) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
