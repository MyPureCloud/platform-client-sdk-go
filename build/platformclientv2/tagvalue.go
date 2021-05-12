package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Tagvalue
type Tagvalue struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name - The workspace tag name.
	Name *string `json:"name,omitempty"`


	// InUse
	InUse *bool `json:"inUse,omitempty"`


	// Acl
	Acl *[]string `json:"acl,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

// String returns a JSON representation of the model
func (o *Tagvalue) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
