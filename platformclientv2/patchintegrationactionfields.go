package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Patchintegrationactionfields
type Patchintegrationactionfields struct { 
	// IntegrationAction - Reference to the Integration Action to be used when integrationAction type is qualified
	IntegrationAction *Patchintegrationaction `json:"integrationAction,omitempty"`


	// RequestMappings - Collection of Request Mappings to use
	RequestMappings *[]Requestmapping `json:"requestMappings,omitempty"`

}

// String returns a JSON representation of the model
func (o *Patchintegrationactionfields) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
