package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Scimphonenumber - Defines a SCIM phone number.
type Scimphonenumber struct { 
	// Value - The phone number in E.164 or tel URI format, for example, tel:+nnnnnnnn; ext=xxxxx.
	Value *string `json:"value,omitempty"`


	// VarType - The type of phone number.
	VarType *string `json:"type,omitempty"`


	// Primary - Indicates whether the phone number is the primary phone number.
	Primary *bool `json:"primary,omitempty"`

}

// String returns a JSON representation of the model
func (o *Scimphonenumber) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
