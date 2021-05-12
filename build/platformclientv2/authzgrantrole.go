package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Authzgrantrole
type Authzgrantrole struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// Description
	Description *string `json:"description,omitempty"`


	// Policies
	Policies *[]Authzgrantpolicy `json:"policies,omitempty"`


	// VarDefault
	VarDefault *bool `json:"default,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

// String returns a JSON representation of the model
func (o *Authzgrantrole) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
