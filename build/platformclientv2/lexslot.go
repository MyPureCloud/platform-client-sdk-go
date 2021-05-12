package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Lexslot
type Lexslot struct { 
	// Name - The slot name
	Name *string `json:"name,omitempty"`


	// Description - The slot description
	Description *string `json:"description,omitempty"`


	// VarType - The slot type
	VarType *string `json:"type,omitempty"`


	// Priority - The priority of the slot
	Priority *int `json:"priority,omitempty"`

}

// String returns a JSON representation of the model
func (o *Lexslot) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
