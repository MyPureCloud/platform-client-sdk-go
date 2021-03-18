package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Resourceconditionvalue
type Resourceconditionvalue struct { 
	// VarType
	VarType *string `json:"type,omitempty"`


	// Value
	Value *string `json:"value,omitempty"`

}

// String returns a JSON representation of the model
func (o *Resourceconditionvalue) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
