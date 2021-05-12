package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Schedulegenerationmessage
type Schedulegenerationmessage struct { 
	// VarType - The type of the message
	VarType *string `json:"type,omitempty"`


	// Arguments - The arguments describing the message
	Arguments *[]Schedulermessageargument `json:"arguments,omitempty"`

}

// String returns a JSON representation of the model
func (o *Schedulegenerationmessage) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
