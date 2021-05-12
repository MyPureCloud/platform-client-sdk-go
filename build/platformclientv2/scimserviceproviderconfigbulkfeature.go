package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Scimserviceproviderconfigbulkfeature - Defines a \"bulk\" request in the SCIM service provider's configuration.
type Scimserviceproviderconfigbulkfeature struct { 
	// Supported - Indicates whether configuration options are supported.
	Supported *bool `json:"supported,omitempty"`


	// MaxOperations - The maximum number of operations for each bulk request.
	MaxOperations *int `json:"maxOperations,omitempty"`


	// MaxPayloadSize - The maximum payload size.
	MaxPayloadSize *int `json:"maxPayloadSize,omitempty"`

}

// String returns a JSON representation of the model
func (o *Scimserviceproviderconfigbulkfeature) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
