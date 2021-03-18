package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Auditentityreference
type Auditentityreference struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// SelfUri
	SelfUri *string `json:"selfUri,omitempty"`


	// VarType
	VarType *string `json:"type,omitempty"`


	// Action
	Action *string `json:"action,omitempty"`

}

// String returns a JSON representation of the model
func (o *Auditentityreference) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
