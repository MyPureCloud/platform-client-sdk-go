package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Shifttradematchviolation
type Shifttradematchviolation struct { 
	// VarType - The type of constraint violation
	VarType *string `json:"type,omitempty"`


	// Params - Clarifying user params for constructing helpful error messages
	Params *map[string]string `json:"params,omitempty"`

}

// String returns a JSON representation of the model
func (o *Shifttradematchviolation) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
