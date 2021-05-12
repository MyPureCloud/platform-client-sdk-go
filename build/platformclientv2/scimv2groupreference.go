package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Scimv2groupreference - Defines a reference to SCIM groups.
type Scimv2groupreference struct { 
	// VarType - The type of SCIM resource.
	VarType *string `json:"type,omitempty"`


	// Value - The ID of the group member. Can be \"userId\" or \"groupId\".
	Value *string `json:"value,omitempty"`


	// Ref - The reference URI of the SCIM resource.
	Ref *string `json:"$ref,omitempty"`

}

// String returns a JSON representation of the model
func (o *Scimv2groupreference) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
