package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Fieldconfig
type Fieldconfig struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// EntityType
	EntityType *string `json:"entityType,omitempty"`


	// State
	State *string `json:"state,omitempty"`


	// Sections
	Sections *[]Section `json:"sections,omitempty"`


	// Version
	Version *string `json:"version,omitempty"`


	// SchemaVersion
	SchemaVersion *string `json:"schemaVersion,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

// String returns a JSON representation of the model
func (o *Fieldconfig) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
