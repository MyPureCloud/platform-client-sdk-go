package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Namedentitydefinition
type Namedentitydefinition struct { 
	// Name - The name of the entity.
	Name *string `json:"name,omitempty"`


	// VarType - The name of the entity type.
	VarType *string `json:"type,omitempty"`

}

// String returns a JSON representation of the model
func (o *Namedentitydefinition) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
