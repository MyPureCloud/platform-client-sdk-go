package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Dependency
type Dependency struct { 
	// Id - The dependency identifier
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// Version
	Version *string `json:"version,omitempty"`


	// VarType
	VarType *string `json:"type,omitempty"`


	// Deleted
	Deleted *bool `json:"deleted,omitempty"`


	// Updated
	Updated *bool `json:"updated,omitempty"`


	// StateUnknown
	StateUnknown *bool `json:"stateUnknown,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

// String returns a JSON representation of the model
func (o *Dependency) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
