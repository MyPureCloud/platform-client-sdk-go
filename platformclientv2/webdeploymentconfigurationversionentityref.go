package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Webdeploymentconfigurationversionentityref
type Webdeploymentconfigurationversionentityref struct { 
	// Id - The configuration version ID
	Id *string `json:"id,omitempty"`


	// Name - The configuration version name
	Name *string `json:"name,omitempty"`


	// SelfUri
	SelfUri *string `json:"selfUri,omitempty"`


	// Version - The version of the configuration
	Version *string `json:"version,omitempty"`

}

// String returns a JSON representation of the model
func (o *Webdeploymentconfigurationversionentityref) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
