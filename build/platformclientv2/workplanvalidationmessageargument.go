package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Workplanvalidationmessageargument
type Workplanvalidationmessageargument struct { 
	// VarType - The type of the argument associated with violation messages
	VarType *string `json:"type,omitempty"`


	// Value - The value of the argument
	Value *string `json:"value,omitempty"`

}

// String returns a JSON representation of the model
func (o *Workplanvalidationmessageargument) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
