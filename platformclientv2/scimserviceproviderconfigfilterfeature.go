package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Scimserviceproviderconfigfilterfeature - Defines a \"filter\" request in the SCIM service provider's configuration.
type Scimserviceproviderconfigfilterfeature struct { 
	// Supported - Indicates whether configuration options are supported.
	Supported *bool `json:"supported,omitempty"`


	// MaxResults - The maximum number of results returned from a filtered query.
	MaxResults *int `json:"maxResults,omitempty"`

}

// String returns a JSON representation of the model
func (o *Scimserviceproviderconfigfilterfeature) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
