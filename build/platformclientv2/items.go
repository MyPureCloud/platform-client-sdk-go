package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Items
type Items struct { 
	// VarType
	VarType *string `json:"type,omitempty"`


	// Pattern
	Pattern *string `json:"pattern,omitempty"`

}

// String returns a JSON representation of the model
func (o *Items) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
