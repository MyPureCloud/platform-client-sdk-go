package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Callbackidentifier
type Callbackidentifier struct { 
	// VarType - The type of the associated callback participant
	VarType *string `json:"type,omitempty"`


	// Id - The identifier of the callback
	Id *string `json:"id,omitempty"`

}

// String returns a JSON representation of the model
func (o *Callbackidentifier) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
