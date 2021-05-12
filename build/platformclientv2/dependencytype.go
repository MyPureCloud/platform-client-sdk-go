package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Dependencytype
type Dependencytype struct { 
	// Id - The dependency type identifier
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// Versioned
	Versioned *bool `json:"versioned,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

// String returns a JSON representation of the model
func (o *Dependencytype) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
