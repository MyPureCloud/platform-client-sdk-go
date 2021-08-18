package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Scimserviceproviderconfigsimplefeature - Defines a request in the SCIM service provider's configuration.
type Scimserviceproviderconfigsimplefeature struct { 
	// Supported - Indicates whether configuration options are supported.
	Supported *bool `json:"supported,omitempty"`

}

func (u *Scimserviceproviderconfigsimplefeature) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Scimserviceproviderconfigsimplefeature

	

	return json.Marshal(&struct { 
		Supported *bool `json:"supported,omitempty"`
		*Alias
	}{ 
		Supported: u.Supported,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Scimserviceproviderconfigsimplefeature) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
