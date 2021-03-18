package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Subjectdivisiongrants
type Subjectdivisiongrants struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// Divisions
	Divisions *[]Division `json:"divisions,omitempty"`


	// VarType
	VarType *string `json:"type,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

// String returns a JSON representation of the model
func (o *Subjectdivisiongrants) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
