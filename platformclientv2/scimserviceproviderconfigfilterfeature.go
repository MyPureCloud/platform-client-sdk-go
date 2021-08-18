package platformclientv2
import (
	"github.com/leekchan/timeutil"
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

func (u *Scimserviceproviderconfigfilterfeature) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Scimserviceproviderconfigfilterfeature

	

	return json.Marshal(&struct { 
		Supported *bool `json:"supported,omitempty"`
		
		MaxResults *int `json:"maxResults,omitempty"`
		*Alias
	}{ 
		Supported: u.Supported,
		
		MaxResults: u.MaxResults,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Scimserviceproviderconfigfilterfeature) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
